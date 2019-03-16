package control

import (
	"../asset"
	"../util"
	"fmt"
	"github.com/nsf/termbox-go"
	"io/ioutil"
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
}

func NewKeyReader() *KeyReader {

	rd := KeyReader{}
	rd.eventChan = make(chan termbox.Event)
	rd.endChan = make(chan struct{})
	rd.currentView = MainMenu
	return &rd

}

/*
This function takes user input into channel.
This function will be called when program starts.
*/

func (rd *KeyReader) Control() {

	err := termbox.Init()
	util.CheckErr(err)
	defer termbox.Close()

	go func() {
		for {
			rd.eventChan <- termbox.PollEvent()
		}
	}()

	rd.menu()

	<-rd.endChan
	close(rd.eventChan)
}

func (rd *KeyReader) refresh() {

	termbox.Clear(asset.ColorEmptyCell, asset.ColorEmptyCell)

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

	err := termbox.Flush()
	util.CheckErr(err)

	for {
		rd.event = <-rd.eventChan

		if rd.event.Type == termbox.EventKey {
			return
		}
	}
}

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
			close(rd.endChan)
			return
		}

	}

}

func (rd *KeyReader) selectMap() {

	rd.currentView = Select

	for {

		rd.refresh()

		switch {
		case rd.event.Key == termbox.KeyEsc:
			rd.currentView = MainMenu
			return
		case rd.event.Key == termbox.KeyArrowUp:
		case rd.event.Key == termbox.KeyArrowDown:
		}

	}

}

func (rd *KeyReader) controlGame() {
	for {
	}
}

func (rd *KeyReader) showMapList() {

	mapList := []string{"[mapList]"}
	files, err := ioutil.ReadDir("./maps")
	util.CheckErr(err)
	for n, file := range files {
		mapList = append(mapList, fmt.Sprintf("%d. %s", n+1, file.Name()))
	}

	rd.printf(5, 3, mapList)

}
