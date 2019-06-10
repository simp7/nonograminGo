package control

import (
	"../asset"
	"../model"
	"../util"
	"github.com/nsf/termbox-go"
)

type View uint8

const (
	MainMenu View = iota
	Select
	InGame
	Result
	Create
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
	case InGame:
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

	rd.currentView = InGame

	correctMap := model.NewNonomap(data)
	playerMap := correctMap.EmptyMap()

	rd.pt = util.NewPlaytime()
	hProblem, vProblem, xpos, ypos := playerMap.CreateProblemFormat()

	for {

		rd.showMap(hProblem, vProblem, xpos, ypos)

		switch {
		case rd.event.Key == termbox.KeyArrowUp:
		case rd.event.Key == termbox.KeyArrowDown:
		case rd.event.Key == termbox.KeyArrowLeft:
		case rd.event.Key == termbox.KeyArrowRight:
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

func (rd *KeyReader) showMap(hProblem []string, vProblem []string, xpos int, ypos int) {
	err := termbox.Clear(asset.ColorEmptyCell, asset.ColorEmptyCell)
	util.CheckErr(err)

	rd.printf(xpos, 0, vProblem)
	rd.printf(0, ypos, hProblem)

	err = termbox.Flush()
	util.CheckErr(err)

	for {
		rd.event = <-rd.eventChan

		if rd.event.Type == termbox.EventKey {
			return
		}
	}
}

func (rd *KeyReader) isComplete() bool {
	return false
}
