package control

import (
	"../asset"
	"../model"
	"../util"
	"github.com/nsf/termbox-go"
)

type View uint8
type Signal uint8

const (
	MainMenu View = iota
	Select
	Result
	Create
)

const (
	Move Signal = iota
	Erase
	Check
	Fill
	//Wrongcheck
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
This function refresh current display because of player's input or time passed
This function will be called when player strokes key or time passed.
*/

func (rd *KeyReader) refresh() {

	err := termbox.Clear(asset.ColorEmptyCell, asset.ColorEmptyCell)
	util.CheckErr(err)

	switch rd.currentView {
	case MainMenu:
		rd.printf(5, 5, asset.StringMainMenu)
	case Select:
		rd.showMapList()
	case Result:
		rd.printf(5, 5, asset.StringResult)
	case Create:
	}

	err = termbox.Flush()
	util.CheckErr(err)

	for {
		rd.event = <-rd.eventChan

		if rd.event.Type == termbox.EventKey {
			return
		}
	} // This loop keeps display until player press key.
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

	rd.currentView = MainMenu

	for {

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

	rd.currentView = Select

	for {

		rd.refresh()

		switch {
		case rd.event.Key == termbox.KeyEsc:
			rd.currentView = MainMenu
			return
		case rd.event.Key == termbox.KeyArrowRight:
		case rd.event.Key == termbox.KeyArrowLeft:
		case rd.event.Ch >= '1' && rd.event.Ch <= '9':
			rd.inGame(rd.fm.GetMapDataByNumber(int(rd.event.Ch - '1')))
		}

	}

}

func (rd *KeyReader) controlGame() {

	for {
	}

}

/*
This function shows the list of the map
This function will be called when refreshing display while being in the select mode
*/

func (rd *KeyReader) showMapList() {

	mapList := []string{"[mapList]   [<-Prev | Next->]"}
	mapList = append(mapList, rd.fm.GetMapList()...)

	rd.printf(5, 3, mapList)

}

func (rd *KeyReader) inGame(data string) {

	correctMap := model.NewNonomap(data)
	playerMap := correctMap.EmptyMap()

	rd.pt = util.NewPlaytime()
	hProblem, vProblem, xProblemPos, yProblemPos := playerMap.CreateProblemFormat()

	xpos, ypos := xProblemPos, yProblemPos+1
	rd.showProblem(hProblem, vProblem, xProblemPos, yProblemPos)

	for {

		rd.showMap(xpos, ypos, Move)

		err := termbox.Flush()
		util.CheckErr(err)

		for {
			rd.event = <-rd.eventChan

			if rd.event.Type == termbox.EventKey {
				break
			}
		}

		switch {
		case rd.event.Key == termbox.KeyArrowUp:
			if ypos-1 >= yProblemPos+1 {
				rd.showMap(xpos, ypos, Erase)
				ypos--
				rd.showMap(xpos, ypos, Move)
			}
		case rd.event.Key == termbox.KeyArrowDown:
			if ypos+1 < yProblemPos+1+playerMap.GetHeight() {
				rd.showMap(xpos, ypos, Erase)
				ypos++
				rd.showMap(xpos, ypos, Move)
			}
		case rd.event.Key == termbox.KeyArrowLeft:
			if xpos-1 >= xProblemPos {
				rd.showMap(xpos, ypos, Erase)
				xpos--
				rd.showMap(xpos, ypos, Move)
			}
		case rd.event.Key == termbox.KeyArrowRight:
			if xpos+1 < xProblemPos+playerMap.GetWidth() {
				rd.showMap(xpos, ypos, Erase)
				xpos++
				rd.showMap(xpos, ypos, Move)
			}
		case rd.event.Key == termbox.KeySpace:
			if rd.isComplete() {
				return
			}

		case rd.event.Ch == 'x' || rd.event.Ch == 'X':

		case rd.event.Key == termbox.KeyEsc:
			rd.pt.EndWitoutResult()
			rd.currentView = Select
			return
		}

	}
}

func (rd *KeyReader) showProblem(hProblem []string, vProblem []string, xpos int, ypos int) {

	err := termbox.Clear(asset.ColorEmptyCell, asset.ColorEmptyCell)
	util.CheckErr(err)

	rd.printf(xpos, 1, vProblem)
	rd.printf(0, ypos+1, hProblem)

	err = termbox.Flush()
	util.CheckErr(err)

}

func (rd *KeyReader) showMap(xpos int, ypos int, signal Signal) {

	switch signal {
	case Move:
		termbox.SetCell(xpos, ypos, '+', asset.ColorFilledCell, asset.ColorEmptyCell)
	case Erase:
		termbox.SetCell(xpos, ypos, '■', asset.ColorEmptyCell, asset.ColorEmptyCell)
	case Check:
		termbox.SetCell(xpos, ypos, 'X', asset.ColorCheckedCell, asset.ColorCheckedCell)
	case Fill:
		termbox.SetCell(xpos, ypos, '■', asset.ColorFilledCell, asset.ColorFilledCell)
	}

}

func (rd *KeyReader) isComplete() bool {
	return false
}
