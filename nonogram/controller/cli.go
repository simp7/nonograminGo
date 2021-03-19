package controller

import (
	"github.com/nsf/termbox-go"
	"github.com/simp7/nonograminGo/nonogram"
	"github.com/simp7/nonograminGo/nonogram/direction"
	"github.com/simp7/nonograminGo/nonogram/fileManager"
	"github.com/simp7/nonograminGo/nonogram/nonomap"
	"github.com/simp7/nonograminGo/nonogram/player"
	"github.com/simp7/nonograminGo/nonogram/setting"
	"github.com/simp7/nonograminGo/nonogram/signal"
	"github.com/simp7/nonograminGo/util"
	"github.com/simp7/times/gadget"
	"github.com/simp7/times/gadget/stopwatch"
	"strconv"
	"sync"
)

type View uint8
type Signal uint8

const (
	MainMenu View = iota
	Select
	Help
	Credit
)

type cli struct {
	eventChan   chan termbox.Event
	endChan     chan struct{}
	currentView View
	event       termbox.Event
	fm          nonogram.FileManager
	timer       gadget.Stopwatch
	locker      sync.Mutex
	*setting.Setting
}

func CLI() nonogram.Controller {

	cc := new(cli)
	cc.eventChan = make(chan termbox.Event)
	cc.endChan = make(chan struct{})
	cc.currentView = MainMenu
	cc.fm = fileManager.New()
	cc.Setting = setting.Get()
	cc.timer = stopwatch.Standard

	return cc

}

/*
This function takes player's input into channel.
This function will be called when program starts.
*/

func (cc *cli) Start() {

	err := termbox.Init()
	util.CheckErr(err)
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
This function wait until player press some keys.
This function would be called when key input is needed.
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
This function refresh current display because of player's input or time passed
This function will be called when player strokes key or time passed.
*/

func (cc *cli) refresh() {

	cc.redraw(func() {
		switch cc.currentView {
		case MainMenu:
			cc.printStandard(cc.MainMenu())
		case Select:
			cc.showMapList()
		case Help:
			cc.printStandard(cc.GetHelp())
		case Credit:
			cc.printStandard(cc.GetCredit())
		}
	})

	cc.pressKeyToContinue()

}

/*
This function prints a list of strings line by line.
This function will be called when display refreshed
*/

func (cc *cli) println(x int, y int, texts []string) {

	temp := x

	for _, msg := range texts {

		for _, ch := range msg {
			termbox.SetCell(x, y, ch, cc.Char, cc.Empty)
			x++
		}

		x = temp
		y++

	}

}

func (cc *cli) printStandard(texts []string) {
	cc.println(cc.DefaultX, cc.DefaultY, texts)
}

/*
This function listens player's input in main menu.
This function will be called when player enters main menu.
*/

func (cc *cli) menu() {

	for {

		cc.currentView = MainMenu
		cc.refresh()

		switch {
		case cc.event.Ch == '1':
			cc.selectMap()
		case cc.event.Ch == '2':
			cc.createNonomapInfo()
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
			cc.fm.NextList()
		case cc.event.Key == termbox.KeyArrowLeft:
			cc.fm.PrevList()
		case cc.event.Ch >= '0' && cc.event.Ch <= '9':
			nonomapData, ok := cc.fm.GetMapDataByNumber(int(cc.event.Ch - '0'))
			if !ok {
				continue
			} else {
				cc.inGame(nonomapData)
			}
		}

	}

}

/*
This function shows the list of the map
This function will be called when refreshing display while being in the select mode
*/

func (cc *cli) showMapList() {

	mapList := make([]string, len(cc.GetSelectHeader()))
	copy(mapList, cc.GetSelectHeader())
	mapList[0] += cc.fm.GetOrder()

	mapList = append(mapList, cc.fm.GetMapList()...)

	cc.printStandard(mapList)

}

/*
This function shows the map current player plays and change its appearance when player press key.
This function will be called when player select map.
*/

func (cc *cli) inGame(correctMap nonogram.Map) {

	util.CheckErr(termbox.Clear(cc.Empty, cc.Empty))

	remainedCell := correctMap.FilledTotal()
	wrongCell := 0

	hProblem, vProblem, xProblemPos, yProblemPos := correctMap.CreateProblemFormat()
	cc.showProblem(hProblem, vProblem, xProblemPos, yProblemPos)

	p := player.New(xProblemPos, yProblemPos, correctMap.GetWidth(), correctMap.GetHeight())
	p.SetMap(signal.Cursor)

	cc.showHeader()

	go cc.timer.Start()

	for {

		err := termbox.Flush()
		util.CheckErr(err)

		cc.pressKeyToContinue()

		switch {

		case cc.event.Key == termbox.KeyArrowUp:
			p.Move(direction.Up)
		case cc.event.Key == termbox.KeyArrowDown:
			p.Move(direction.Down)
		case cc.event.Key == termbox.KeyArrowLeft:
			p.Move(direction.Left)
		case cc.event.Key == termbox.KeyArrowRight:
			p.Move(direction.Right)
		case cc.event.Key == termbox.KeySpace || cc.event.Ch == 'z' || cc.event.Ch == 'Z':

			if p.GetMapSignal() == signal.Empty {

				if correctMap.ShouldFilled(p.RealPos()) {
					p.Toggle(signal.Fill)
					remainedCell--

					if remainedCell == 0 { //Enter when p complete the game
						p.SetMap(signal.Fill)
						cc.showResult(wrongCell)
						return
					}

				} else {
					p.Toggle(signal.Wrong)
					wrongCell++
				}

			}

		case cc.event.Ch == 'x' || cc.event.Ch == 'X':
			if p.GetMapSignal() == signal.Empty {
				p.Toggle(signal.Check)
			} else if p.GetMapSignal() == signal.Check {
				p.Toggle(signal.Empty)
			}

		case cc.event.Key == termbox.KeyEsc:
			cc.timer.Stop()
			return
		}

	}

}

func (cc *cli) showProblem(hProblem []string, vProblem []string, xPos int, yPos int) {

	cc.redraw(func() {
		cc.println(xPos, 1, vProblem)
		cc.println(0, yPos+1, hProblem)
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

	result[3] += cc.fm.GetCurrentMapName()
	result[4] += cc.timer.Stop()
	result[5] += strconv.Itoa(wrong)

	cc.locker.Lock()

	cc.println(0, 0, []string{cc.Complete()})
	util.CheckErr(termbox.Flush())

	cc.pressKeyToContinue()
	cc.locker.Unlock()

	cc.redraw(func() { cc.printStandard(result) })

	cc.pressKeyToContinue()

}

/*
	This function receive user's key input to create name of nonogram map in create mode.
	This function will be called when player enter the create mode from main menu.
*/

func (cc *cli) createNonomapInfo() {

	width, height := 0, 0
	var err error
	criteria := nonomap.New()

	mapName := cc.stringReader(cc.RequestMapName())
	if mapName == "" {
		return
	}

	mapWidth := cc.stringReader(cc.RequestWidth())
	for {
		if mapWidth == "" {
			return
		} else {
			width, err = strconv.Atoi(mapWidth)
			util.CheckErr(err)
			if width <= criteria.WidthLimit() {
				break
			}
			mapWidth = cc.stringReader(cc.SizeError() + strconv.Itoa(criteria.WidthLimit()))
		}
	}

	mapHeight := cc.stringReader(cc.RequestHeight())
	for {
		if mapHeight == "" {
			return
		} else {
			height, err = strconv.Atoi(mapHeight)
			util.CheckErr(err)
			if height <= criteria.HeightLimit() {
				break
			}
			mapHeight = cc.stringReader(cc.SizeError() + strconv.Itoa(criteria.HeightLimit()))
		}
	}

	cc.inCreate(mapName, width, height)

}

/*
	This function gets string value from player.
	This function will be called when player creates map so configures properties of map.
*/

func (cc *cli) stringReader(header string) (result string) {

	result = ""
	resultByte := make([]rune, cc.NameMax)
	n := 0

	cc.redraw(func() { cc.printStandard([]string{header}) })

	for {
		cc.pressKeyToContinue()

		cc.redraw(func() {
			cc.printStandard([]string{header})

			if n < cc.NameMax {
				if header == cc.RequestMapName() {
					if cc.event.Ch != 0 {
						resultByte[n] = cc.event.Ch
						n++
					} else if cc.event.Key == termbox.KeySpace {
						resultByte[n] = ' '
						n++
					}
				} else if cc.event.Ch >= '0' && cc.event.Ch <= '9' {
					resultByte[n] = cc.event.Ch
					n++
				}
			}

			if (cc.event.Key == termbox.KeyBackspace || cc.event.Key == termbox.KeyBackspace2 || cc.event.Key == termbox.KeyDelete) && n > 0 {
				n--
			}

			result = ""
			for i := 0; i < n; i++ {
				result += string(resultByte[i])
			}

			cc.println(cc.DefaultX, cc.DefaultY+2, []string{result})

		})

		if cc.event.Key == termbox.KeyEnter {
			return
		} else if cc.event.Key == termbox.KeyEsc {
			result = ""
			return
		}

	}

}

/*
	This function shows player's current map in create mode and receive player's key input.
	This function will be called when player finish writing name of nonomap that player would create.
*/

func (cc *cli) inCreate(mapName string, width int, height int) {

	cc.redraw(func() { cc.println(1, 0, []string{mapName}) })

	p := player.New(cc.DefaultX, cc.DefaultY, width, height)
	p.SetMap(signal.Cursor)

	for {
		err := termbox.Flush()
		util.CheckErr(err)

		cc.pressKeyToContinue()

		switch {

		case cc.event.Key == termbox.KeyArrowUp:
			p.Move(direction.Up)
		case cc.event.Key == termbox.KeyArrowDown:
			p.Move(direction.Down)
		case cc.event.Key == termbox.KeyArrowLeft:
			p.Move(direction.Left)
		case cc.event.Key == termbox.KeyArrowRight:
			p.Move(direction.Right)
		case cc.event.Key == termbox.KeySpace || cc.event.Ch == 'z' || cc.event.Ch == 'Z':
			if p.GetMapSignal() == signal.Empty {
				p.Toggle(signal.Fill)
			} else if p.GetMapSignal() == signal.Fill {
				p.Toggle(signal.Empty)
			}
		case cc.event.Ch == 'x' || cc.event.Ch == 'X':
			if p.GetMapSignal() == signal.Empty {
				p.Toggle(signal.Check)
			} else if p.GetMapSignal() == signal.Check {
				p.Toggle(signal.Empty)
			}
		case cc.event.Key == termbox.KeyEsc:
			return
		case cc.event.Key == termbox.KeyEnter:
			cc.fm.CreateMap(mapName, width, height, p.FinishCreating())
			cc.fm.RefreshMapList()
			return
		}

	}
}

/*
	This function shows time passed in game.
	This function will be called when player enter the game.
	This function should be called as goroutine and should finish when player finish the game.
*/

func (cc *cli) showHeader() {

	mapName := cc.fm.GetCurrentMapName()

	cc.timer.Add(func(current string) {
		cc.println(cc.DefaultX, 0, []string{mapName + cc.BlankBetweenMapNameAndTimer() + current})
		util.CheckErr(termbox.Flush())
	})

}

/*
	This function erase existing things in display and draw things in function.
	This function will be called when display has to be cleared.
*/

func (cc *cli) redraw(function func()) {

	util.CheckErr(termbox.Clear(cc.Empty, cc.Empty))
	function()
	util.CheckErr(termbox.Flush())

}
