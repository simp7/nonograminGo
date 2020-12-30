package control

import (
	"github.com/nsf/termbox-go"
	"github.com/simp7/nonograminGo/asset"
	"github.com/simp7/nonograminGo/model"
	"github.com/simp7/nonograminGo/util"
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

type Controller interface {
	Start()
}

type cliController struct {
	eventChan   chan termbox.Event
	endChan     chan struct{}
	currentView View
	event       termbox.Event
	fm          FileManager
	timer       util.Timer
	locker      sync.Mutex
}

func NewCliController() Controller {

	rd := cliController{}
	rd.eventChan = make(chan termbox.Event)
	rd.endChan = make(chan struct{})
	rd.currentView = MainMenu
	rd.fm = NewFileManager()
	return &rd

}

/*
This function takes player's input into channel.
This function will be called when program starts.
*/

func (cc *cliController) Start() {

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

func (cc *cliController) pressKeyToContinue() {

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

func (cc *cliController) refresh() {

	redraw(func() {
		switch cc.currentView {
		case MainMenu:
			cc.println(asset.NumberDefaultX, asset.NumberDefaultY, asset.StringMainMenu)
		case Select:
			cc.showMapList()
		case Help:
			cc.println(asset.NumberDefaultX, asset.NumberDefaultY, asset.StringHelp)
		case Credit:
			cc.println(asset.NumberDefaultX, asset.NumberDefaultY, asset.StringCredit)
		}
	})

	cc.pressKeyToContinue()

}

/*
This function prints a list of strings line by line.
This function will be called when display refreshed
*/

func (cc *cliController) println(x int, y int, texts []string) {

	temp := x

	for _, msg := range texts {

		for _, ch := range msg {
			termbox.SetCell(x, y, ch, asset.ColorText, asset.ColorEmptyCell)
			x++
		}

		x = temp
		y++

	}

}

/*
This function listens player's input in main menu.
This function will be called when player enters main menu.
*/

func (cc *cliController) menu() {

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

func (cc *cliController) selectMap() {

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
			nonomapData := cc.fm.GetMapDataByNumber(int(cc.event.Ch - '0'))
			if nonomapData == asset.StringMsgFileNotExist {
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

func (cc *cliController) showMapList() {

	mapList := make([]string, len(asset.StringSelectHeader))
	copy(mapList, asset.StringSelectHeader)
	mapList[0] += cc.fm.GetOrder()

	mapList = append(mapList, cc.fm.GetMapList()...)

	cc.println(asset.NumberDefaultX, asset.NumberDefaultY, mapList)

}

/*
This function shows the map current player plays and change its appearance when player press key.
This function will be called when player select map.
*/

func (cc *cliController) inGame(data string) {

	util.CheckErr(termbox.Clear(asset.ColorEmptyCell, asset.ColorEmptyCell))
	correctMap := model.NewNonomap(data)

	remainedCell := correctMap.TotalCells()
	wrongCell := 0

	hProblem, vProblem, xProblemPos, yProblemPos := correctMap.CreateProblemFormat()
	cc.showProblem(hProblem, vProblem, xProblemPos, yProblemPos)

	cc.timer = util.NewPlaytime()

	player := model.NewPlayer(xProblemPos, yProblemPos, correctMap.GetWidth(), correctMap.GetHeight())
	player.SetMap(model.Cursor)

	go cc.showHeader()

	for {

		err := termbox.Flush()
		util.CheckErr(err)

		cc.pressKeyToContinue()

		switch {

		case cc.event.Key == termbox.KeyArrowUp:
			player.Move(model.Up)
		case cc.event.Key == termbox.KeyArrowDown:
			player.Move(model.Down)
		case cc.event.Key == termbox.KeyArrowLeft:
			player.Move(model.Left)
		case cc.event.Key == termbox.KeyArrowRight:
			player.Move(model.Right)
		case cc.event.Key == termbox.KeySpace || cc.event.Ch == 'z' || cc.event.Ch == 'Z':

			if player.GetMapSignal() == model.Empty {
				if correctMap.CompareValidity(player.RealPos()) {
					player.SetMap(model.CursorFilled)
					player.SetMapSignal(model.Fill)
					remainedCell--

					if remainedCell == 0 { //Enter when player complete the game
						player.SetMap(model.Fill)
						cc.showResult(wrongCell)
						return
					}

				} else {
					player.SetMap(model.CursorWrong)
					player.SetMapSignal(model.Wrong)
					wrongCell++
				}

			}

		case cc.event.Ch == 'x' || cc.event.Ch == 'X':
			if player.GetMapSignal() == model.Empty {
				player.SetMap(model.CursorChecked)
				player.SetMapSignal(model.Check)
			} else if player.GetMapSignal() == model.Check {
				player.SetMap(model.Cursor)
				player.SetMapSignal(model.Empty)
			}

		case cc.event.Key == termbox.KeyEsc:
			cc.timer.End()
			return
		}

	}

}

func (cc *cliController) showProblem(hProblem []string, vProblem []string, xPos int, yPos int) {

	redraw(func() {
		cc.println(xPos, 1, vProblem)
		cc.println(0, yPos+1, hProblem)
	})

}

/*
	This function shows total result in game.
	This function will be called when player finally solve the problem and after seeing the whole answer picture.
*/

func (cc *cliController) showResult(wrong int) {

	resultFormat := asset.StringResult
	result := make([]string, len(resultFormat))
	copy(result, resultFormat)

	result[4] += cc.fm.GetCurrentMapName()
	result[5] += cc.timer.GetResult()
	result[6] += strconv.Itoa(wrong)

	cc.locker.Lock()

	cc.println(0, 0, asset.StringComplete)
	util.CheckErr(termbox.Flush())

	cc.pressKeyToContinue()
	cc.locker.Unlock()

	redraw(func() { cc.println(asset.NumberDefaultX, asset.NumberDefaultY, result) })

	cc.pressKeyToContinue()

}

/*
	This function receive user's key input to create name of nonogram map in create mode.
	This function will be called when player enter the create mode from main menu.
*/

func (cc *cliController) createNonomapInfo() {

	width, height := 0, 0
	var err error

	mapName := cc.stringReader(asset.StringHeaderMapName)
	if mapName == "" {
		return
	}

	mapWidth := cc.stringReader(asset.StringHeaderWidth)
	for {
		if mapWidth == "" {
			return
		} else {
			width, err = strconv.Atoi(mapWidth)
			util.CheckErr(err)
			if width <= asset.NumberWidthMax {
				break
			}
			mapWidth = cc.stringReader(asset.StringHeaderSizeError + strconv.Itoa(asset.NumberWidthMax))
		}
	}

	mapHeight := cc.stringReader(asset.StringHeaderHeight)
	for {
		if mapHeight == "" {
			return
		} else {
			height, err = strconv.Atoi(mapHeight)
			util.CheckErr(err)
			if height <= asset.NumberHeightMax {
				break
			}
			mapHeight = cc.stringReader(asset.StringHeaderSizeError + strconv.Itoa(asset.NumberHeightMax))
		}
	}

	cc.inCreate(mapName, width, height)

}

/*
	This function gets string value from player.
	This function will be called when player creates map so configures properties of map.
*/

func (cc *cliController) stringReader(header string) (result string) {

	result = ""
	resultByte := make([]rune, asset.NumberNameMax)
	n := 0

	redraw(func() { cc.println(asset.NumberDefaultX, asset.NumberDefaultY, []string{header}) })

	for {
		cc.pressKeyToContinue()

		redraw(func() {
			cc.println(asset.NumberDefaultX, asset.NumberDefaultY, []string{header})

			if n < asset.NumberNameMax {
				if header == asset.StringHeaderMapName {
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

			cc.println(asset.NumberDefaultX, asset.NumberDefaultY+2, []string{result})

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

func (cc *cliController) inCreate(mapName string, width int, height int) {

	redraw(func() { cc.println(1, 0, []string{mapName}) })

	player := model.NewPlayer(asset.NumberDefaultX, asset.NumberDefaultY, width, height)
	player.SetMap(model.Cursor)

	for {
		err := termbox.Flush()
		util.CheckErr(err)

		cc.pressKeyToContinue()

		switch {

		case cc.event.Key == termbox.KeyArrowUp:
			player.Move(model.Up)
		case cc.event.Key == termbox.KeyArrowDown:
			player.Move(model.Down)
		case cc.event.Key == termbox.KeyArrowLeft:
			player.Move(model.Left)
		case cc.event.Key == termbox.KeyArrowRight:
			player.Move(model.Right)
		case cc.event.Key == termbox.KeySpace || cc.event.Ch == 'z' || cc.event.Ch == 'Z':
			if player.GetMapSignal() == model.Empty {
				player.SetMap(model.CursorFilled)
				player.SetMapSignal(model.Fill)
			} else if player.GetMapSignal() == model.Fill {
				player.SetMap(model.Cursor)
				player.SetMapSignal(model.Empty)
			}
		case cc.event.Ch == 'x' || cc.event.Ch == 'X':
			if player.GetMapSignal() == model.Empty {
				player.SetMap(model.CursorChecked)
				player.SetMapSignal(model.Check)
			} else if player.GetMapSignal() == model.Check {
				player.SetMap(model.Cursor)
				player.SetMapSignal(model.Empty)
			}
		case cc.event.Key == termbox.KeyEsc:
			return
		case cc.event.Key == termbox.KeyEnter:
			cc.fm.CreateMap(mapName, width, height, player.ConvertToBitMap())
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

func (cc *cliController) showHeader() {

	mapName := cc.fm.GetCurrentMapName()

	cc.locker.Lock()
	defer cc.locker.Unlock()

	cc.timer.Do(func(current string) {
		cc.println(asset.NumberDefaultX, 0, []string{mapName + asset.StringBlankBetweenMapAndTimer + current})
		util.CheckErr(termbox.Flush())
	})

}

/*
	This function erase existing things in display and draw things in function.
	This function will be called when display has to be cleared.
*/

func redraw(function func()) {

	util.CheckErr(termbox.Clear(asset.ColorEmptyCell, asset.ColorEmptyCell))
	function()
	util.CheckErr(termbox.Flush())

}
