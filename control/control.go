package control

import (
	"../asset"
	"../model"
	"../util"
	"github.com/nsf/termbox-go"
	"strconv"
)

type View uint8
type Signal uint8

const (
	MainMenu View = iota
	Select
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
			rd.currentView = Credit
			rd.refresh()
		case rd.event.Ch == '4' || rd.event.Key == termbox.KeyEsc:
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

	mapList := asset.StringSelectHeader
	mapList = append(mapList, rd.fm.GetMapList()...)

	rd.printf(asset.NumberDefaultX, asset.NumberDefaultY, mapList)

}

/*
This function shows the map current player plays and change its appearence when player press key.
This function will be called when player select map.
*/

func (rd *KeyReader) inGame(data string) {

	correctMap := model.NewNonomap(data)

	remainedCell := correctMap.TotalCells()
	wrongCell := 0

	err := termbox.Clear(asset.ColorEmptyCell, asset.ColorEmptyCell)
	util.CheckErr(err)

	rd.pt = util.NewPlaytime()

	hProblem, vProblem, xProblemPos, yProblemPos := correctMap.CreateProblemFormat()
	rd.showProblem(hProblem, vProblem, xProblemPos, yProblemPos)

	playermap := initializeMap(correctMap.GetWidth(), correctMap.GetHeight())
	rd.printf(1, 0, []string{rd.fm.GetCurrentMapName()})

	for n := range playermap {
		for m := range playermap[n] {
			rd.setMap((2*m)+xProblemPos, n+yProblemPos+1, Empty)
		}
	}

	xpos, ypos := xProblemPos, yProblemPos+1
	rd.setMap(xpos, ypos, Cursor)

	for {

		getRealpos := func() (realxpos int, realypos int) {
			realxpos, realypos = (xpos-xProblemPos)/2, ypos-yProblemPos-1
			return
		}

		getMapSignal := func() Signal {
			realxpos, realypos := getRealpos()
			return playermap[realypos][realxpos]
		}

		setMapSignal := func(xpos int, ypos int, signal Signal) {
			realxpos, realypos := getRealpos()
			playermap[realypos][realxpos] = signal
		}

		moveCursor := func(condition bool, function func()) {
			if condition {
				rd.setMap(xpos, ypos, getMapSignal())
				function()
				rd.setCursor(xpos, ypos, getMapSignal())
			}
		}

		err := termbox.Flush()
		util.CheckErr(err)

		rd.pressKeyToContinue()

		switch {

		case rd.event.Key == termbox.KeyArrowUp:
			moveCursor(ypos-1 >= yProblemPos+1, func() { ypos-- })

		case rd.event.Key == termbox.KeyArrowDown:
			moveCursor(ypos+1 < yProblemPos+1+correctMap.GetHeight(), func() { ypos++ })

		case rd.event.Key == termbox.KeyArrowLeft:
			moveCursor(xpos-2 >= xProblemPos, func() { xpos -= 2 })

		case rd.event.Key == termbox.KeyArrowRight:
			moveCursor(xpos+2 < xProblemPos+(2*correctMap.GetWidth()), func() { xpos += 2 })

		case rd.event.Key == termbox.KeySpace:
			if getMapSignal() == Check || getMapSignal() == Fill || getMapSignal() == Wrong {
				continue
			}

			if correctMap.CompareValidity(getRealpos()) {
				rd.setMap(xpos, ypos, CursorFilled)
				setMapSignal(xpos, ypos, Fill)
				remainedCell--

				if remainedCell == 0 { //Enter when player complete the game
					redrow(func() {
						rd.showAnswer(playermap)
					})
					rd.pressKeyToContinue()
					rd.showResult(wrongCell)
					return
				}

			} else {
				rd.setMap(xpos, ypos, CursorWrong)
				setMapSignal(xpos, ypos, Wrong)
				wrongCell++
			}

		case rd.event.Ch == 'x' || rd.event.Ch == 'X':
			if getMapSignal() == Empty {
				rd.setMap(xpos, ypos, CursorChecked)
				setMapSignal(xpos, ypos, Check)
			} else if getMapSignal() == Check {
				rd.setMap(xpos, ypos, Cursor)
				setMapSignal(xpos, ypos, Empty)
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
	This function set the cell in game with signal.
	This function would be called when player press key in game.
*/

func (rd *KeyReader) setMap(xpos int, ypos int, signal Signal) {

	switch signal {
	case Cursor:
		termbox.SetCell(xpos, ypos, '(', asset.ColorFilledCell, asset.ColorEmptyCell)
		termbox.SetCell(xpos+1, ypos, ')', asset.ColorFilledCell, asset.ColorEmptyCell)
	case Empty:
		termbox.SetCell(xpos, ypos, ' ', asset.ColorEmptyCell, asset.ColorEmptyCell)
		termbox.SetCell(xpos+1, ypos, ' ', asset.ColorEmptyCell, asset.ColorEmptyCell)
	case Check:
		termbox.SetCell(xpos, ypos, '>', asset.ColorCheckedCell, asset.ColorEmptyCell)
		termbox.SetCell(xpos+1, ypos, '<', asset.ColorCheckedCell, asset.ColorEmptyCell)
	case Fill:
		termbox.SetCell(xpos, ypos, ' ', asset.ColorFilledCell, asset.ColorFilledCell)
		termbox.SetCell(xpos+1, ypos, ' ', asset.ColorFilledCell, asset.ColorFilledCell)
	case Wrong:
		termbox.SetCell(xpos, ypos, '>', asset.ColorWrongCell, asset.ColorEmptyCell)
		termbox.SetCell(xpos+1, ypos, '<', asset.ColorWrongCell, asset.ColorEmptyCell)
	case CursorFilled:
		termbox.SetCell(xpos, ypos, '(', asset.ColorEmptyCell, asset.ColorFilledCell)
		termbox.SetCell(xpos+1, ypos, ')', asset.ColorEmptyCell, asset.ColorFilledCell)
	case CursorWrong:
		termbox.SetCell(xpos, ypos, '(', asset.ColorWrongCell, asset.ColorEmptyCell)
		termbox.SetCell(xpos+1, ypos, ')', asset.ColorWrongCell, asset.ColorEmptyCell)
	case CursorChecked:
		termbox.SetCell(xpos, ypos, '(', asset.ColorCheckedCell, asset.ColorEmptyCell)
		termbox.SetCell(xpos+1, ypos, ')', asset.ColorCheckedCell, asset.ColorEmptyCell)
	}

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

	redrow(func() {
		rd.printf(asset.NumberDefaultX, asset.NumberDefaultY, result)
	})

	rd.pressKeyToContinue()

}

/*
	This function shows whole answer picture that player solve.
	This function will be called when player finally solve the problem.
*/

func (rd *KeyReader) showAnswer(playermap [][]Signal) {

	rd.printf(asset.NumberDefaultX, asset.NumberDefaultY, []string{"You Complete Me!"})
	for n := range playermap {
		for m, v := range playermap[n] {
			if v == Fill {
				rd.setMap((2*m)+asset.NumberDefaultX, n+asset.NumberDefaultY+3, Fill)
			} else {
				rd.setMap((2*m)+asset.NumberDefaultX, n+asset.NumberDefaultY+3, Empty)
			}
		}
	}

}

/*
	This function receive user's key input to create name of nonogram map in create mode.
	This function will be called when player enter the create mode from main menu.
*/
func (rd *KeyReader) createNonomapInfo() {

	mapName := rd.stringReader(asset.StringHeaderMapname)
	if mapName == "" {
		return
	}
	mapWidth := rd.stringReader(asset.StringHeaderWidth)
	if mapWidth == "" {
		return
	}
	mapHeight := rd.stringReader(asset.StringHeaderHeight)
	if mapHeight == "" {
		return
	}

	width, err := strconv.Atoi(mapWidth)
	util.CheckErr(err)
	height, err := strconv.Atoi(mapHeight)
	util.CheckErr(err)
	rd.inCreate(mapName, width, height)

}
func (rd *KeyReader) stringReader(header string) (result string) {

	result = ""
	resultByte := make([]rune, asset.NumberNameMax)
	n := 0

	redrow(func() {
		rd.printf(asset.NumberDefaultX, asset.NumberDefaultY, []string{header})
	})

	for {
		rd.pressKeyToContinue()

		redrow(func() {
			rd.printf(asset.NumberDefaultX, asset.NumberDefaultY, []string{header})

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
			if (rd.event.Key == termbox.KeyBackspace || rd.event.Key == termbox.KeyBackspace2 || rd.event.Key == termbox.KeyDelete) && n > 0 {
				resultByte[n] = 0
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

	answermap := make([][]bool, height)
	for n := range answermap {
		answermap[n] = make([]bool, width)
	}

	playermap := initializeMap(width, height)

	redrow(func() {
		rd.printf(1, 0, []string{mapName})
	})

	for n := range playermap {
		for m := range playermap[n] {
			rd.setMap((2*m)+asset.NumberDefaultX, n+asset.NumberDefaultY, Empty)
			answermap[n][m] = false
		}
	}

	xpos, ypos := asset.NumberDefaultX, asset.NumberDefaultY
	rd.setMap(xpos, ypos, Cursor)

	for {
		realxpos, realypos := (xpos-asset.NumberDefaultX)/2, ypos-asset.NumberDefaultY

		err := termbox.Flush()
		util.CheckErr(err)

		rd.pressKeyToContinue()

		switch {

		case rd.event.Key == termbox.KeyArrowUp:
			if ypos-1 >= asset.NumberDefaultY {
				rd.setMap(xpos, ypos, playermap[realypos][realxpos])
				ypos--
				rd.setMap(xpos, ypos, Cursor)
			}

		case rd.event.Key == termbox.KeyArrowDown:
			if ypos+1 < asset.NumberDefaultY+height {
				rd.setMap(xpos, ypos, playermap[realypos][realxpos])
				ypos++
				rd.setMap(xpos, ypos, Cursor)
			}
		case rd.event.Key == termbox.KeyArrowLeft:
			if xpos-2 >= asset.NumberDefaultX {
				rd.setMap(xpos, ypos, playermap[realypos][realxpos])
				xpos -= 2
				rd.setMap(xpos, ypos, Cursor)
			}
		case rd.event.Key == termbox.KeyArrowRight:
			if xpos+2 < asset.NumberDefaultX+(2*width) {
				rd.setMap(xpos, ypos, playermap[realypos][realxpos])
				xpos += 2
				rd.setMap(xpos, ypos, Cursor)
			}
		case rd.event.Key == termbox.KeySpace:
			if playermap[realypos][realxpos] == Empty {
				playermap[realypos][realxpos] = Fill
				answermap[realypos][realxpos] = true
			} else if playermap[realypos][realxpos] == Fill {
				playermap[realypos][realxpos] = Empty
				answermap[realypos][realxpos] = false
			}
		case rd.event.Ch == 'x' || rd.event.Ch == 'X':
			if playermap[realypos][realxpos] == Empty {
				playermap[realypos][realxpos] = Check
			} else if playermap[realypos][realxpos] == Check {
				playermap[realypos][realxpos] = Empty
			}
		case rd.event.Key == termbox.KeyEsc:
			return
		case rd.event.Key == termbox.KeyEnter:
			rd.fm.CreateMap(mapName, width, height, answermap)
			rd.fm.RefreshMapList()
			return
		}

	}
}

/*
	This function initialize current player's map data to compare the map with answer.
	The result will be used in showing player's current map.
	The function will be called when player enter the game.
*/

func (rd *KeyReader) setCursor(xpos int, ypos int, CellState Signal) {
	if CellState == Fill {
		rd.setMap(xpos, ypos, CursorFilled)
	} else if CellState == Check {
		rd.setMap(xpos, ypos, CursorChecked)
	} else if CellState == Wrong {
		rd.setMap(xpos, ypos, CursorWrong)
	} else {
		rd.setMap(xpos, ypos, Cursor)
	}
}

func initializeMap(width int, height int) (emptyMap [][]Signal) {

	emptyMap = make([][]Signal, height)

	for n := range emptyMap {
		emptyMap[n] = make([]Signal, width)
		for m := range emptyMap[n] {
			emptyMap[n][m] = Empty
		}

	}

	return

}

/*
	This function erase existing things in display and drow things in function.
	This function will be called when display has to be cleared.
*/

func redrow(function func()) {

	err := termbox.Clear(asset.ColorEmptyCell, asset.ColorEmptyCell)
	util.CheckErr(err)

	function()

	err = termbox.Flush()
	util.CheckErr(err)
}
