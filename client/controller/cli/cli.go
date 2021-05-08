package cli

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/simp7/nonograminGo/client"
	"github.com/simp7/nonograminGo/file"
	"github.com/simp7/nonograminGo/nonogram"
	"github.com/simp7/times/gadget"
	"github.com/simp7/times/gadget/stopwatch"
	"io"
	"log"
	"strconv"
	"sync"
	"unicode"
)

type View uint8

const (
	MainMenu View = iota
	Select
	Help
	Credit
)

type cli struct {
	eventChan   chan termbox.Event
	endChan     chan struct{}
	nonomap     nonogram.Map
	fileSystem  file.System
	currentView View
	event       termbox.Event
	timer       gadget.Stopwatch
	locker      sync.Mutex
	mapList     file.MapList
	*Config
}

/*
	Controller returns nonogram.Controller that runs in Controller
*/

func Controller(fileSystem file.System, formatter file.Formatter, mapPrototype nonogram.Map) client.Controller {

	var err error
	cc := new(cli)

	cc.eventChan = make(chan termbox.Event)
	cc.endChan = make(chan struct{})

	cc.Config, err = InitSetting(fileSystem, formatter)
	checkErr(err)

	cc.nonomap = mapPrototype
	cc.fileSystem = fileSystem
	cc.currentView = MainMenu
	cc.mapList = fileSystem.Maps()
	cc.timer = stopwatch.Standard

	return cc

}

/*
Start() takes player's input into channel.
Start() function will be called when program starts.
*/

func (cc *cli) Start() {

	err := termbox.Init()
	checkErr(err)
	defer termbox.Close()

	go func() {
		for {
			select {
			case cc.eventChan <- termbox.PollEvent():
			case <-cc.endChan:
				return
			}
		}
	}()

	cc.menu()

	close(cc.eventChan)

}

/*
pressKeyToContinue() wait until player press some keys.
pressKeyToContinue() would be called when key input is needed.
*/

func (cc *cli) pressKeyToContinue() {

	for {

		cc.event = <-cc.eventChan

		if cc.event.Type == termbox.EventKey {
			return
		}

	}

}

/*
refresh() refreshes current display because of player's input or time passed.
refresh() will be called when player strokes key or time passed.
*/

func (cc *cli) refresh() {

	cc.redraw(func() {
		switch cc.currentView {
		case MainMenu:
			cc.printStandard(cc.MainMenu()...)
		case Select:
			cc.showMapList()
		case Help:
			cc.printStandard(cc.GetHelp()...)
		case Credit:
			cc.printStandard(cc.GetCredit()...)
		}
	})

	cc.pressKeyToContinue()

}

/*
isCJK() determines character if it is CJK(Chinese-Japanese-Korean).
isCJK() is only called in print() because printing CJK needs two cells.
*/

func isCJK(char rune) bool {
	return unicode.In(char, unicode.Hangul, unicode.Han, unicode.Hiragana, unicode.Katakana)
}

/*
print() prints a list of strings line by line.
print() will be called when display refreshed
*/

func (cc *cli) print(position Pos, texts ...string) {

	y := 0

	for _, msg := range texts {
		cc.println(position.Move(0, y), msg)
		y++
	}

}

func (cc *cli) println(position Pos, text string) {

	x := position.X

	for _, ch := range text {
		termbox.SetCell(x, position.Y, ch, cc.Char, cc.Empty)
		if isCJK(ch) {
			x++
		}
		x++
	}

}

/*
printStandard() is simplified version of print().
The position of text is fixed in defaultX and defaultY
*/

func (cc *cli) printStandard(texts ...string) {
	cc.print(cc.DefaultPos, texts...)
}

/*
menu() listens player's input in main menu.
menu() will be called when player enters main menu.
*/

func (cc *cli) menu() {

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

/*
This function listens player's input in map-select
This function will be called when player enters map-select.
*/

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
			name, ok := cc.mapList.GetMapName(int(cc.event.Ch - '0'))
			if !ok {
				continue
			} else {
				cc.inGame(cc.loadMap(name))
			}
		}

	}

}

func (cc *cli) loadMap(name string) nonogram.Map {

	mapData := cc.nonomap

	s, err := cc.fileSystem.Map(name, mapData.Formatter())
	checkErr(err)

	err = s.Load(&mapData)
	checkErr(err)

	return mapData

}

/*
This function shows the list of the map
This function will be called when refreshing display while being in the select mode
*/

func (cc *cli) showMapList() {

	list := make([]string, len(cc.SelectHeader()))
	copy(list, cc.SelectHeader())
	list[0] += fmt.Sprintf("(%d/%d)", cc.mapList.CurrentPage(), cc.mapList.LastPage())

	list = append(list, cc.mapList.Current()...)

	cc.printStandard(list...)

}

/*
This function shows the map current player plays and change its appearance when player press key.
This function will be called when player select map.
*/

func (cc *cli) inGame(correctMap nonogram.Map) {

	checkErr(termbox.Clear(cc.Empty, cc.Empty))

	remainedCell := correctMap.FilledTotal()
	wrongCell := 0

	problem := correctMap.CreateProblem()
	cc.showProblem(problem)

	p := Player(cc.Config.Color, Pos{problem.Horizontal().Max(), problem.Vertical().Max()}, correctMap.GetWidth(), correctMap.GetHeight())
	p.SetCell(Cursor)

	cc.showHeader()

	go cc.timer.Start()

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
			cc.timer.Stop()
			return
		}

	}

}

func (cc *cli) showProblem(problem nonogram.Problem) {

	cc.redraw(func() {

		verticalPos := Pos{problem.Horizontal().Max(), 1}
		horizontalPos := Pos{0, problem.Vertical().Max() + 1}

		cc.print(verticalPos, problem.Vertical().Get()...)
		cc.print(horizontalPos, problem.Horizontal().Get()...)

	})

}

/*
	This function shows total result in game.
	This function will be called when player finally solve the problem and after seeing the whole answer picture.
*/

func (cc *cli) showResult(wrong int) {

	resultFormat := cc.GetResult()
	result := make([]string, len(resultFormat))
	copy(result, resultFormat)

	cc.locker.Lock()
	result[3] += cc.mapList.GetCachedMapName()
	result[4] += cc.timer.Stop()
	result[5] += strconv.Itoa(wrong)
	cc.locker.Unlock()

	cc.print(Pos{0, 0}, cc.Complete())
	checkErr(termbox.Flush())

	cc.pressKeyToContinue()

	cc.redraw(func() { cc.printStandard(result...) })

	cc.pressKeyToContinue()

}

/*
	This function receive user's key input to create name of nonogram map in create mode.
	This function will be called when player enter the create mode from main menu.
*/

func (cc *cli) createNonomapSkeleton() {

	width, height := 0, 0
	var err error
	criteria := cc.nonomap
	header := cc.RequestMapName()

	mapName := cc.stringReader(header, cc.NameMax)
	if mapName == "" {
		return
	}

	header = cc.RequestWidth()
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
		header = cc.SizeError() + strconv.Itoa(criteria.WidthLimit())

	}

	header = cc.RequestHeight()
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
		header = cc.SizeError() + strconv.Itoa(criteria.HeightLimit())

	}

	cc.inCreate(mapName, width, height)

}

/*
	This function gets string value from player.
	This function will be called when player creates map so configures properties of map.
*/

func (cc *cli) stringReader(header string, maxLen int) string {

	resultByte := make([]rune, 0)

	writeChar := func(ch rune) {
		resultByte = append(resultByte, ch)
	}

	placeholder := func() {

		cc.printStandard(header)
		if len(resultByte) < maxLen {
			cc.print(cc.DefaultPos.Move(len(resultByte), 2), "_")
		}

		if cc.DefaultPos.X > 0 {
			cc.print(cc.DefaultPos.Move(-1, 2), "[")
			cc.print(cc.DefaultPos.Move(maxLen, 2), "]")
		}

	}

	cc.redraw(func() {
		placeholder()
	})

	for {

		cc.pressKeyToContinue()

		cc.redraw(func() {

			defer func() {
				cc.print(cc.DefaultPos.Move(0, 2), string(resultByte))
				placeholder()
			}()

			if (cc.event.Key == termbox.KeyBackspace || cc.event.Key == termbox.KeyBackspace2 || cc.event.Key == termbox.KeyDelete) && len(resultByte) > 0 {
				resultByte = resultByte[:len(resultByte)-1]
			}

			if len(resultByte) == maxLen {
				return
			}

			if header == cc.RequestMapName() {
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

/*
	This function shows player's current map in create mode and receive player's key input.
	This function will be called when player finish writing name of nonomap that player would create.
*/

func (cc *cli) inCreate(mapName string, width int, height int) {

	cc.redraw(func() { cc.print(Pos{1, 0}, mapName) })

	p := Player(cc.Config.Color, cc.DefaultPos, width, height)
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
			cc.saveMap(mapName, p.FinishCreating(cc.nonomap))
			checkErr(cc.mapList.Refresh())
			return

		}

	}

}

func (cc *cli) saveMap(name string, mapData nonogram.Map) {

	mapSaver, err := cc.fileSystem.Map(name, mapData.Formatter())
	checkErr(err)

	checkErr(mapSaver.Save(mapData))

}

/*
	This function shows time passed in game.
	This function will be called when player enter the game.
	This function should be called as goroutine and should finish when player finish the game.
*/

func (cc *cli) showHeader() {

	mapName := cc.mapList.GetCachedMapName()

	cc.timer.Add(func(current string) {
		cc.print(Pos{cc.DefaultPos.X, 0}, mapName+cc.BlankBetweenMapNameAndTimer()+current)
		checkErr(termbox.Flush())
	})

}

/*
	This function erase existing things in display and draw things in function.
	This function will be called when display has to be cleared.
*/

func (cc *cli) redraw(function func()) {

	checkErr(termbox.Clear(cc.Empty, cc.Empty))
	function()
	checkErr(termbox.Flush())

}

func checkErr(e error) {

	if e == nil || e == io.EOF {
		return
	}

	if termbox.IsInit {
		termbox.Close()
	}
	log.Fatal(e)

}
