package control

import (
	"../asset"
	"../model"
	"../util"
	"github.com/nsf/termbox-go"
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

const (
	Cursor Signal = iota
	Empty
	Check
	Fill
	Wrong
	CursorFilled
	CursorChecked
	CursorWrong
)

type KeyReader struct {
	eventChan   chan termbox.Event
	endChan     chan struct{}
	currentView View
	event       termbox.Event
	fm          *FileManager
	pt          *util.Playtime
	locker      sync.Mutex
}

func NewKeyReader() *KeyReader {

	rd := KeyReader{}
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

func (rd *KeyReader) Control() {

	err := termbox.Init()
	util.CheckErr(err)
	defer termbox.Close()

	go func() {
		for {
			select {
			case rd.eventChan <- termbox.PollEvent():
			case <-rd.endChan:
				return
			}
		}
	}()

	rd.menu()

	close(rd.eventChan)

}

/*
This function wait until player press some keys.
This function would be called when key input is needed.
*/

func (rd *KeyReader) pressKeyToContinue() {

	for {
		rd.event = <-rd.eventChan

		if rd.event.Type == termbox.EventKey {
			return
		}
	}

}

/*
This function refresh current display because of player's input or time passed
This function will be called when player strokes key or time passed.
*/

func (rd *KeyReader) refresh() {

	redrow(func() {
		switch rd.currentView {
		case MainMenu:
			rd.printf(asset.NumberDefaultX, asset.NumberDefaultY, asset.StringMainMenu)
		case Select:
			rd.showMapList()
		case Help:
			rd.printf(asset.NumberDefaultX, asset.NumberDefaultY, asset.StringHelp)
		case Credit:
			rd.printf(asset.NumberDefaultX, asset.NumberDefaultY, asset.StringCredit)
		}
	})

	rd.pressKeyToContinue()

}

/*
This function prints a list of strings line by line.
This function will be called when display refreshed
*/

func (rd *KeyReader) printf(x int, y int, msgs []string) {

	temp := x

	for _, msg := range msgs {

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

func (rd *KeyReader) menu() {

	for {

		rd.currentView = MainMenu
		rd.refresh()

		switch {
		case rd.event.Ch == '1':
			rd.selectMap()
		case rd.event.Ch == '2':
			rd.createNonomapInfo()
		case rd.event.Ch == '3':
			rd.currentView = Help
			rd.refresh()
		case rd.event.Ch == '4':
			rd.currentView = Credit
			rd.refresh()
		case rd.event.Ch == '5' || rd.event.Key == termbox.KeyEsc:
			return
		}

	}

}

/*
This function listens player's input in map-select
This function will be called when player enters map-select.
*/

func (rd *KeyReader) selectMap() {

	for {

		rd.currentView = Select
		rd.refresh()

		switch {
		case rd.event.Key == termbox.KeyEsc:
			return
		case rd.event.Key == termbox.KeyArrowRight:
			rd.fm.NextList()
		case rd.event.Key == termbox.KeyArrowLeft:
			rd.fm.PrevList()
		case rd.event.Ch >= '0' && rd.event.Ch <= '9':
			nonomapData := rd.fm.GetMapDataByNumber(int(rd.event.Ch - '0'))
			if nonomapData == asset.StringMsgFileNotExist {
				continue
			} else {
				rd.inGame(nonomapData)
			}
		}

	}

}

/*
This function shows the list of the map
This function will be called when refreshing display while being in the select mode
*/

func (rd *KeyReader) showMapList() {

	mapList := make([]string, len(asset.StringSelectHeader))
	copy(mapList, asset.StringSelectHeader)
	mapList[0] += rd.fm.GetOrder()

	mapList = append(mapList, rd.fm.GetMapList()...)

	rd.printf(asset.NumberDefaultX, asset.NumberDefaultY, mapList)

}

/*
This function shows the map current player plays and change its appearence when player press key.
This function will be called when player select map.
*/

func (rd *KeyReader) inGame(data string) {

	util.CheckErr(termbox.Clear(asset.ColorEmptyCell, asset.ColorEmptyCell))
	correctMap := model.NewNonomap(data)

	remainedCell := correctMap.TotalCells()
	wrongCell := 0

	hProblem, vProblem, xProblemPos, yProblemPos := correctMap.CreateProblemFormat()
	rd.showProblem(hProblem, vProblem, xProblemPos, yProblemPos)

	rd.pt = util.NewPlaytime()

	player := model.NewPlayer(xProblemPos, yProblemPos, correctMap.GetWidth(), correctMap.GetHeight())
	player.SetMap(model.Cursor)

	go rd.showTimePassed()

	for {

		err := termbox.Flush()
		util.CheckErr(err)

		rd.pressKeyToContinue()

		switch {

		case rd.event.Key == termbox.KeyArrowUp:
			player.Move(model.Up)
		case rd.event.Key == termbox.KeyArrowDown:
			player.Move(model.Down)
		case rd.event.Key == termbox.KeyArrowLeft:
			player.Move(model.Left)
		case rd.event.Key == termbox.KeyArrowRight:
			player.Move(model.Right)
		case rd.event.Key == termbox.KeySpace || rd.event.Ch == 'z' || rd.event.Ch == 'Z':

			if player.GetMapSignal() == model.Empty {
				if correctMap.CompareValidity(player.GetRealpos()) {
					player.SetMap(model.CursorFilled)
					player.SetMapSignal(model.Fill)
					remainedCell--

					if remainedCell == 0 { //Enter when player complete the game
						player.SetMap(model.Fill)
						rd.showResult(wrongCell)
						return
					}

				} else {
					player.SetMap(model.CursorWrong)
					player.SetMapSignal(model.Wrong)
					wrongCell++
				}

			}

		case rd.event.Ch == 'x' || rd.event.Ch == 'X':
			if player.GetMapSignal() == model.Empty {
				player.SetMap(model.CursorChecked)
				player.SetMapSignal(model.Check)
			} else if player.GetMapSignal() == model.Check {
				player.SetMap(model.Cursor)
				player.SetMapSignal(model.Empty)
			}

		case rd.event.Key == termbox.KeyEsc:
			rd.pt.EndWithoutResult()
			return
		}

	}

}

func (rd *KeyReader) showProblem(hProblem []string, vProblem []string, xpos int, ypos int) {

	redrow(func() {
		rd.printf(xpos, 1, vProblem)
		rd.printf(0, ypos+1, hProblem)
	})

}

/*
	This function shows total result in game.
	This function will be called when player finally solve the problem and after seeing the whole answer picture.
*/

func (rd *KeyReader) showResult(wrong int) {

	resultFormat := asset.StringResult
	result := make([]string, len(resultFormat))
	copy(result, resultFormat)

	result[4] += rd.fm.GetCurrentMapName()
	result[5] += rd.pt.TimeResult()
	result[6] += strconv.Itoa(wrong)

	rd.locker.Lock()

	rd.printf(0, 0, asset.StringComplete)
	util.CheckErr(termbox.Flush())

	rd.pressKeyToContinue()
	rd.locker.Unlock()

	redrow(func() { rd.printf(asset.NumberDefaultX, asset.NumberDefaultY, result) })

	rd.pressKeyToContinue()

}

/*
	This function receive user's key input to create name of nonogram map in create mode.
	This function will be called when player enter the create mode from main menu.
*/

func (rd *KeyReader) createNonomapInfo() {

	width, height := 0, 0
	var err error

	mapName := rd.stringReader(asset.StringHeaderMapname)
	if mapName == "" {
		return
	}

	mapWidth := rd.stringReader(asset.StringHeaderWidth)
	for {
		if mapWidth == "" {
			return
		} else {
			width, err = strconv.Atoi(mapWidth)
			util.CheckErr(err)
			if width <= asset.NumberWidthMax {
				break
			}
			mapWidth = rd.stringReader(asset.StringHeaderSizeError + strconv.Itoa(asset.NumberWidthMax))
		}
	}

	mapHeight := rd.stringReader(asset.StringHeaderHeight)
	for {
		if mapHeight == "" {
			return
		} else {
			height, err = strconv.Atoi(mapHeight)
			util.CheckErr(err)
			if height <= asset.NumberHeightMax {
				break
			}
			mapHeight = rd.stringReader(asset.StringHeaderSizeError + strconv.Itoa(asset.NumberHeightMax))
		}
	}

	rd.inCreate(mapName, width, height)

}

/*
	This function gets string value from player.
	This function will be called when player creates map so configures properties of map.
*/

func (rd *KeyReader) stringReader(header string) (result string) {

	result = ""
	resultByte := make([]rune, asset.NumberNameMax)
	n := 0

	redrow(func() { rd.printf(asset.NumberDefaultX, asset.NumberDefaultY, []string{header}) })

	for {
		rd.pressKeyToContinue()

		redrow(func() {
			rd.printf(asset.NumberDefaultX, asset.NumberDefaultY, []string{header})

			if n < asset.NumberNameMax {
				if header == asset.StringHeaderMapname {
					if rd.event.Ch != 0 {
						resultByte[n] = rd.event.Ch
						n++
					} else if rd.event.Key == termbox.KeySpace {
						resultByte[n] = ' '
						n++
					}
				} else if rd.event.Ch >= '0' && rd.event.Ch <= '9' {
					resultByte[n] = rd.event.Ch
					n++
				}
			}

			if (rd.event.Key == termbox.KeyBackspace || rd.event.Key == termbox.KeyBackspace2 || rd.event.Key == termbox.KeyDelete) && n > 0 {
				n--
			}

			result = ""
			for i := 0; i < n; i++ {
				result += string(resultByte[i])
			}

			rd.printf(asset.NumberDefaultX, asset.NumberDefaultY+2, []string{result})

		})

		if rd.event.Key == termbox.KeyEnter {
			return
		} else if rd.event.Key == termbox.KeyEsc {
			result = ""
			return
		}

	}

}

/*
	This function shows player's current map in create mode and receive player's key input.
	This function will be called when player finish writing name of nonomap that player would create.
*/

func (rd *KeyReader) inCreate(mapName string, width int, height int) {

	redrow(func() { rd.printf(1, 0, []string{mapName}) })

	player := model.NewPlayer(asset.NumberDefaultX, asset.NumberDefaultY, width, height)
	player.SetMap(model.Cursor)

	for {
		err := termbox.Flush()
		util.CheckErr(err)

		rd.pressKeyToContinue()

		switch {

		case rd.event.Key == termbox.KeyArrowUp:
			player.Move(model.Up)
		case rd.event.Key == termbox.KeyArrowDown:
			player.Move(model.Down)
		case rd.event.Key == termbox.KeyArrowLeft:
			player.Move(model.Left)
		case rd.event.Key == termbox.KeyArrowRight:
			player.Move(model.Right)
		case rd.event.Key == termbox.KeySpace || rd.event.Ch == 'z' || rd.event.Ch == 'Z':
			if player.GetMapSignal() == model.Empty {
				player.SetMap(model.CursorFilled)
				player.SetMapSignal(model.Fill)
			} else if player.GetMapSignal() == model.Fill {
				player.SetMap(model.Cursor)
				player.SetMapSignal(model.Empty)
			}
		case rd.event.Ch == 'x' || rd.event.Ch == 'X':
			if player.GetMapSignal() == model.Empty {
				player.SetMap(model.CursorChecked)
				player.SetMapSignal(model.Check)
			} else if player.GetMapSignal() == model.Check {
				player.SetMap(model.Cursor)
				player.SetMapSignal(model.Empty)
			}
		case rd.event.Key == termbox.KeyEsc:
			return
		case rd.event.Key == termbox.KeyEnter:
			rd.fm.CreateMap(mapName, width, height, player.ConvertToBitMap())
			rd.fm.RefreshMapList()
			return
		}

	}
}

/*
	This function shows time passed in game.
	This function will be called when player enter the game.
	This function should be called as goroutine and should finish when player finish the game.
*/

func (rd *KeyReader) showTimePassed() {

	mapname := rd.fm.GetCurrentMapName()

	rd.locker.Lock()
	defer rd.locker.Unlock()

	for {
		select {
		case current := <-rd.pt.Clock:
			rd.printf(asset.NumberDefaultX, 0, []string{mapname + asset.StringBlankBetweenMapAndTimer + current})
			util.CheckErr(termbox.Flush())
		case <-rd.pt.Stop:
			return
		}
	}

}

/*
	This function erase existing things in display and drow things in function.
	This function will be called when display has to be cleared.
*/

func redrow(function func()) {

	util.CheckErr(termbox.Clear(asset.ColorEmptyCell, asset.ColorEmptyCell))

	function()

	util.CheckErr(termbox.Flush())
}
