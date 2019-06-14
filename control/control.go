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
	Create
)

const (
	Cursor Signal = iota
	Empty
	Check
	Fill
	Wrong
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
			rd.printf(5, 5, asset.StringMainMenu)
		case Select:
			rd.showMapList()
		case Create:
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
		case rd.event.Ch == '3':
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
		case rd.event.Key == termbox.KeyArrowLeft:
		case rd.event.Ch >= '1' && rd.event.Ch <= '9':
			rd.inGame(rd.fm.GetMapDataByNumber(int(rd.event.Ch - '1')))
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

	rd.printf(5, 3, mapList)

}

/*
This function shows the map current player plays and change its appearence when player press key.
This function will be called when player select map.
*/

func (rd *KeyReader) inGame(data string) {

	correctMap := model.NewNonomap(data)

	remainedCell := correctMap.TotalCells()
	wrongCell := 0

	rd.pt = util.NewPlaytime()
	hProblem, vProblem, xProblemPos, yProblemPos := correctMap.CreateProblemFormat()

	xpos, ypos := xProblemPos, yProblemPos+1
	rd.showProblem(hProblem, vProblem, xProblemPos, yProblemPos)

	playermap := initializeMap(correctMap.GetWidth(), correctMap.GetHeight())
	rd.printf(1, 0, []string{rd.fm.GetCurrentMapName()})

	for n := range playermap {
		for m := range playermap[n] {
			rd.setMap(m+xProblemPos, n+yProblemPos+1, Empty)
		}
	}

	rd.setMap(xpos, ypos, Cursor)

	for {

		realxpos, realypos := xpos-xProblemPos, ypos-yProblemPos-1

		err := termbox.Flush()
		util.CheckErr(err)

		rd.pressKeyToContinue()

		switch {

		case rd.event.Key == termbox.KeyArrowUp:
			if ypos-1 >= yProblemPos+1 {
				rd.setMap(xpos, ypos, playermap[realypos][realxpos])
				ypos--
				rd.setMap(xpos, ypos, Cursor)
			}

		case rd.event.Key == termbox.KeyArrowDown:
			if ypos+1 < yProblemPos+1+correctMap.GetHeight() {
				rd.setMap(xpos, ypos, playermap[realypos][realxpos])
				ypos++
				rd.setMap(xpos, ypos, Cursor)
			}

		case rd.event.Key == termbox.KeyArrowLeft:
			if xpos-1 >= xProblemPos {
				rd.setMap(xpos, ypos, playermap[realypos][realxpos])
				xpos--
				rd.setMap(xpos, ypos, Cursor)
			}

		case rd.event.Key == termbox.KeyArrowRight:
			if xpos+1 < xProblemPos+correctMap.GetWidth() {
				rd.setMap(xpos, ypos, playermap[realypos][realxpos])
				xpos++
				rd.setMap(xpos, ypos, Cursor)
			}

		case rd.event.Key == termbox.KeySpace:
			if playermap[realypos][realxpos] == Check || playermap[realypos][realxpos] == Fill {
				continue
			}

			if correctMap.CompareValidity(realxpos, realypos) {
				rd.setMap(xpos, ypos, Fill)
				playermap[realypos][realxpos] = Fill
				remainedCell--
				if remainedCell == 0 {
					redrow(func() {
						rd.printf(1, 0, []string{"You Complete Me!"})
						rd.showAnswer(playermap)
					})
					rd.pressKeyToContinue()
					rd.ShowResult(wrongCell)
					return
				}
			} else {
				rd.setMap(xpos, ypos, Wrong)
				playermap[realypos][realxpos] = Wrong
				wrongCell++
			}

		case rd.event.Ch == 'x' || rd.event.Ch == 'X':
			if playermap[realypos][realxpos] == Empty {
				rd.setMap(xpos, ypos, Check)
				playermap[realypos][realxpos] = Check
			} else if playermap[realypos][realxpos] == Check {
				rd.setMap(xpos, ypos, Empty)
				playermap[realypos][realxpos] = Empty
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
		termbox.SetCell(xpos, ypos, '+', asset.ColorFilledCell, asset.ColorEmptyCell)
	case Empty:
		termbox.SetCell(xpos, ypos, '☐', asset.ColorFilledCell, asset.ColorEmptyCell)
	case Check:
		termbox.SetCell(xpos, ypos, 'x', asset.ColorCheckedCell, asset.ColorEmptyCell)
	case Fill:
		termbox.SetCell(xpos, ypos, '■', asset.ColorFilledCell, asset.ColorEmptyCell)
	case Wrong:
		termbox.SetCell(xpos, ypos, '☐', asset.ColorCheckedCell, asset.ColorEmptyCell)
	}

}

/*
	This function shows total result in game.
	This function will be called when player finally solve the problem and after seeing the whole answer picture.
*/

func (rd *KeyReader) ShowResult(wrong int) {

	resultFormat := asset.StringResult
	result := make([]string, len(resultFormat))
	copy(result, resultFormat)

	result[4] += rd.fm.GetCurrentMapName()
	result[5] += rd.pt.TimeResult()
	result[6] += strconv.Itoa(wrong)

	redrow(func() {
		rd.printf(5, 5, result)
	})

	rd.pressKeyToContinue()

}

/*
	This function shows whole answer picture that player solve.
	This function will be called when player finally solve the problem.
*/

func (rd *KeyReader) showAnswer(playermap [][]Signal) {

	for n := range playermap {
		for m, v := range playermap[n] {
			if v == Fill {
				rd.setMap(m+2, n+2, Fill)
			} else {
				rd.setMap(m+2, n+2, Empty)
			}
		}
	}

}

/*
	This function initialize current player's map data to compare the map with answer.
	The result will be used in showing player's current map.
	The function will be called when player enter the game.
*/

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
