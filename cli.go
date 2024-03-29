package main

import (
	"fmt"
	"github.com/gdamore/tcell/termbox"
	"github.com/simp7/nonogram"
	"github.com/simp7/nonogram/unit"
	"github.com/simp7/times/gadget"
	"github.com/simp7/times/gadget/stopwatch"
	"io"
	"log"
	"strconv"
	"unicode"
)

type view uint8

const (
	MainMenu view = iota
	Select
	Help
	Credit
)

type cli struct {
	eventChan   chan termbox.Event
	endChan     chan struct{}
	currentView view
	core        *nonogram.Core
	event       termbox.Event
	stopwatch   gadget.Stopwatch
	mapList     *mapList
	config      Config
}

//Controller returns nonogram.Controller that runs in Controller
func Controller(core *nonogram.Core) *cli {

	cc := new(cli)

	cc.eventChan = make(chan termbox.Event)
	cc.endChan = make(chan struct{})

	cc.core = core

	config, err := core.LoadSetting()
	checkErr(err)

	cc.config = AdaptConfig(config)

	cc.currentView = MainMenu
	cc.refreshMapList()
	cc.stopwatch = stopwatch.Standard

	return cc

}

func (cc *cli) Start() {

	err := termbox.Init()
	checkErr(err)
	defer termbox.Close()

	go cc.startEventHandler()

	cc.menu()
	<-cc.endChan

}

func (cc *cli) startEventHandler() {
	for {
		select {
		case cc.eventChan <- termbox.PollEvent():
		case <-cc.endChan:
			close(cc.eventChan)
			return
		}
	}
}

func (cc *cli) pressKeyToContinue() {
	for {
		if cc.event = <-cc.eventChan; cc.event.Type == termbox.EventKey {
			return
		}
	}
}

func (cc *cli) refresh() {

	cc.redraw(func() {
		switch cc.currentView {
		case MainMenu:
			cc.printStandard(cc.config.MainMenu()...)
		case Select:
			cc.showMapList()
		case Help:
			cc.printStandard(cc.config.GetHelp()...)
		case Credit:
			cc.printStandard(cc.config.GetCredit()...)
		}
	})

	cc.pressKeyToContinue()

}

func isCJK(char rune) bool {
	return unicode.In(char, unicode.Hangul, unicode.Han, unicode.Hiragana, unicode.Katakana)
}

func (cc *cli) print(position Pos, texts ...string) {
	for y, msg := range texts {
		cc.println(position.Move(0, y), msg)
	}
}

func (cc *cli) println(position Pos, text string) {

	x := position.X

	for _, ch := range text {
		termbox.SetCell(x, position.Y, ch, cc.config.Char, cc.config.Empty)
		if isCJK(ch) {
			x++
		}
		x++
	}

}

func (cc *cli) printStandard(texts ...string) {
	cc.print(cc.config.DefaultPos, texts...)
}

func (cc *cli) menu() {

	defer close(cc.endChan)

	for {

		cc.currentView = MainMenu
		cc.refresh()

		switch {
		case cc.event.Ch == '1':
			cc.selectMap()
		case cc.event.Ch == '2':
			cc.createNonomapSkeleton()
		case cc.event.Ch == '3':
			cc.currentView = Help
			cc.refresh()
		case cc.event.Ch == '4':
			cc.currentView = Credit
			cc.refresh()
		case cc.event.Ch == '5' || cc.event.Key == termbox.KeyEsc:
			return
		}

	}

}

func (cc *cli) selectMap() {

	for {

		cc.currentView = Select
		cc.refresh()

		switch {
		case cc.event.Key == termbox.KeyEsc:
			return
		case cc.event.Key == termbox.KeyArrowRight:
			cc.mapList.Next()
		case cc.event.Key == termbox.KeyArrowLeft:
			cc.mapList.Prev()
		case cc.event.Ch >= '0' && cc.event.Ch <= '9':
			if name, ok := cc.mapList.GetMapName(int(cc.event.Ch - '0')); ok {
				nonomap, err := cc.core.LoadMap(name)
				checkErr(err)
				cc.inGame(nonomap)
			}
		}

	}

}

func (cc *cli) showMapList() {

	if cc.mapList.IsEmpty() {
		cc.redraw(func() { cc.printStandard(cc.config.FileNotExist()) })
		cc.pressKeyToContinue()
		return
	}

	list := make([]string, len(cc.config.SelectHeader()))
	copy(list, cc.config.SelectHeader())
	list[0] += fmt.Sprintf("(%d/%d)", cc.mapList.CurrentPage(), cc.mapList.LastPage())

	list = append(list, cc.mapList.Current()...)

	cc.printStandard(list...)

}

func (cc *cli) inGame(correctMap unit.Map) {

	termbox.Clear(cc.config.Empty, cc.config.Empty)

	remainedCell := correctMap.FilledTotal()
	wrongCell := 0

	problem := correctMap.CreateProblem()
	cc.showProblem(correctMap)

	p := Player(cc.config.Color, Pos{2 * problem.Horizontal().Max(), problem.Vertical().Max()}, correctMap.Width(), correctMap.Height(), cc.core)
	p.SetCell(Cursor)

	cc.showHeader()

	go cc.stopwatch.Start()

	for {

		err := termbox.Flush()
		checkErr(err)

		cc.pressKeyToContinue()

		switch {

		case cc.event.Key == termbox.KeyArrowUp:
			p.Move(Up)
		case cc.event.Key == termbox.KeyArrowDown:
			p.Move(Down)
		case cc.event.Key == termbox.KeyArrowLeft:
			p.Move(Left)
		case cc.event.Key == termbox.KeyArrowRight:
			p.Move(Right)
		case cc.event.Key == termbox.KeySpace || cc.event.Ch == 'z' || cc.event.Ch == 'Z':

			if p.GetMapSignal() == Empty {

				if correctMap.ShouldFilled(p.RealPos().X, p.RealPos().Y) {
					p.Toggle(Fill)
					remainedCell--

					if remainedCell == 0 { //Enter when p complete the game
						p.SetCell(Fill)
						cc.showResult(wrongCell)
						return
					}

				} else {
					p.Toggle(Wrong)
					wrongCell++
				}

			}

		case cc.event.Ch == 'x' || cc.event.Ch == 'X':
			if p.GetMapSignal() == Empty {
				p.Toggle(Check)
			} else if p.GetMapSignal() == Check {
				p.Toggle(Empty)
			}

		case cc.event.Key == termbox.KeyEsc:
			cc.stopwatch.Stop()
			return
		}

	}

}

func (cc *cli) formatVertical(nonomap unit.Map) []string {

	vertical := nonomap.CreateProblem().Vertical()
	max := vertical.Max()

	problem := make([]string, max)

	for i := max; i > 0; i-- {
		problem[max-i] = ""
		for j := 0; j < nonomap.Width(); j++ {
			currentRow := vertical.Get(j)
			if i > len(currentRow) {
				problem[max-i] += "  "
				continue
			}
			if currentRow[len(currentRow)-i] < 10 {
				problem[max-i] += " "
			}
			problem[max-i] += strconv.Itoa(currentRow[len(currentRow)-i])
		}
	}

	return problem

}

func (cc *cli) formatHorizontal(nonomap unit.Map) []string {

	horizontal := nonomap.CreateProblem().Horizontal()
	max := horizontal.Max()

	problem := make([]string, nonomap.Height())

	for i := 0; i < nonomap.Height(); i++ {
		currentRow := horizontal.Get(i)
		problem[i] = ""
		for j := max; j > 0; j-- {
			if len(currentRow) < j {
				problem[i] += "  "
				continue
			}
			if currentRow[len(currentRow)-j] < 10 {
				problem[i] += " "
			}
			problem[i] += strconv.Itoa(currentRow[len(currentRow)-j])
		}
	}

	return problem

}

func (cc *cli) showProblem(nonomap unit.Map) {

	cc.redraw(func() {

		problem := nonomap.CreateProblem()

		hMax := problem.Horizontal().Max()
		vMax := problem.Vertical().Max()

		verticalPos := Pos{hMax * 2, 1}
		horizontalPos := Pos{0, vMax + 1}

		cc.print(horizontalPos, cc.formatHorizontal(nonomap)...)
		cc.print(verticalPos, cc.formatVertical(nonomap)...)

	})

}

func (cc *cli) showResult(wrong int) {

	resultFormat := cc.config.GetResult()
	result := make([]string, len(resultFormat))
	copy(result, resultFormat)

	result[3] += cc.mapList.GetCachedMapName()
	result[4] += cc.stopwatch.Stop()
	result[5] += strconv.Itoa(wrong)

	cc.print(Pos{0, 0}, cc.config.Complete())
	checkErr(termbox.Flush())

	cc.pressKeyToContinue()
	cc.redraw(func() { cc.printStandard(result...) })
	cc.pressKeyToContinue()

}

func (cc *cli) createNonomapSkeleton() {

	var width, height int
	var err error
	criteria := cc.core.InitMap([][]bool{{true}})
	header := cc.config.RequestMapName()

	mapName := cc.stringReader(header, cc.config.NameMax)
	if mapName == "" {
		return
	}

	header = cc.config.RequestWidth()
	for {

		mapWidth := cc.stringReader(header, 2)
		if mapWidth == "" {
			return
		}

		width, err = strconv.Atoi(mapWidth)
		checkErr(err)

		if width <= criteria.WidthLimit() && width > 0 {
			break
		}
		header = cc.config.SizeError() + strconv.Itoa(criteria.WidthLimit())

	}

	header = cc.config.RequestHeight()
	for {

		mapHeight := cc.stringReader(header, 2)
		if mapHeight == "" {
			return
		}

		height, err = strconv.Atoi(mapHeight)
		checkErr(err)

		if height <= criteria.HeightLimit() && height > 0 {
			break
		}
		header = cc.config.SizeError() + strconv.Itoa(criteria.HeightLimit())

	}

	cc.inCreate(mapName, width, height)

}

func (cc *cli) stringReader(header string, maxLen int) string {

	resultByte := make([]rune, 0)

	writeChar := func(ch rune) {
		resultByte = append(resultByte, ch)
	}

	placeholder := func() {

		defaultPos := cc.config.DefaultPos

		cc.printStandard(header)
		if len(resultByte) < maxLen {
			cc.print(defaultPos.Move(len(resultByte), 2), "_")
		}

		if defaultPos.X > 0 {
			cc.print(defaultPos.Move(-1, 2), "[")
			cc.print(defaultPos.Move(maxLen, 2), "]")
		}

	}

	cc.redraw(func() {
		placeholder()
	})

	for {

		cc.pressKeyToContinue()

		cc.redraw(func() {

			defer func() {
				cc.print(cc.config.DefaultPos.Move(0, 2), string(resultByte))
				placeholder()
			}()

			if (cc.event.Key == termbox.KeyBackspace || cc.event.Key == termbox.KeyBackspace2 || cc.event.Key == termbox.KeyDelete) && len(resultByte) > 0 {
				resultByte = resultByte[:len(resultByte)-1]
			}

			if len(resultByte) == maxLen {
				return
			}

			if header == cc.config.RequestMapName() {
				if cc.event.Ch != 0 {
					writeChar(cc.event.Ch)
				} else if cc.event.Key == termbox.KeySpace {
					writeChar(' ')
				}
			} else if cc.event.Ch >= '0' && cc.event.Ch <= '9' {
				writeChar(cc.event.Ch)
			}

		})

		if cc.event.Key == termbox.KeyEnter {
			return string(resultByte)
		} else if cc.event.Key == termbox.KeyEsc {
			return ""
		}

	}

}

func (cc *cli) inCreate(mapName string, width int, height int) {

	cc.redraw(func() { cc.print(Pos{1, 0}, mapName) })

	p := Player(cc.config.Color, cc.config.DefaultPos, width, height, cc.core)
	p.SetCell(Cursor)

	for {

		err := termbox.Flush()
		checkErr(err)

		cc.pressKeyToContinue()

		switch {

		case cc.event.Key == termbox.KeyArrowUp:
			p.Move(Up)
		case cc.event.Key == termbox.KeyArrowDown:
			p.Move(Down)
		case cc.event.Key == termbox.KeyArrowLeft:
			p.Move(Left)
		case cc.event.Key == termbox.KeyArrowRight:
			p.Move(Right)
		case cc.event.Key == termbox.KeySpace || cc.event.Ch == 'z' || cc.event.Ch == 'Z':
			if p.GetMapSignal() == Empty {
				p.Toggle(Fill)
			} else if p.GetMapSignal() == Fill {
				p.Toggle(Empty)
			}
		case cc.event.Ch == 'x' || cc.event.Ch == 'X':
			if p.GetMapSignal() == Empty {
				p.Toggle(Check)
			} else if p.GetMapSignal() == Check {
				p.Toggle(Empty)
			}
		case cc.event.Key == termbox.KeyEsc:
			return
		case cc.event.Key == termbox.KeyEnter:
			checkErr(cc.core.SaveMap(mapName, p.FinishCreating()))
			cc.refreshMapList()
			return
		}

	}

}

func (cc *cli) refreshMapList() {
	maps, _ := cc.core.Maps()
	cc.mapList = newMapList(maps)
}

func (cc *cli) showHeader() {

	mapName := cc.mapList.GetCachedMapName()

	cc.stopwatch.Add(func(current string) {
		cc.print(Pos{cc.config.DefaultPos.X, 0}, mapName+cc.config.BlankBetweenMapNameAndTimer()+current)
		checkErr(termbox.Flush())
	})

}

func (cc *cli) redraw(function func()) {

	termbox.Clear(cc.config.Empty, cc.config.Empty)
	function()
	checkErr(termbox.Flush())

}

func checkErr(e error) {

	if e == nil || e == io.EOF {
		return
	}

	termbox.Close()
	log.Fatal(e)

}
