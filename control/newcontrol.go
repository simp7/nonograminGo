package control

import (
	"../util"
	"../view"
	"bufio"
	"github.com/nsf/termbox-go"
	"os"
)

type keyReader struct {
	eventChan    chan termbox.Event
	currentEvent termbox.Event
}

func NewKeyReader() *keyReader {

	rd := keyReader{}
	rd.eventChan = make(chan termbox.Event)
	return &rd

}

func (rd *keyReader) poll() {

	rd.eventChan <- termbox.PollEvent()
	rd.currentEvent = <-rd.eventChan

}

func (rd *keyReader) ControlMenu() {

	for {
		rd.poll()
		switch rd.currentEvent.Ch {
		case '1':
			view.showMapList()
			rd.SelectMap()
		case '2':

		case '3':
		case '4':
			return
		}
	}

}

func (rd *keyReader) SelectMap() {

	nameReader := bufio.NewReader(os.Stdin)

	for {
		mapName, _, err := nameReader.ReadLine()
		util.CheckErr(err)
		targetMap := newNonomap(string(mapName))
	}

}

func (rd *keyReader) ControlGame() {
	for {
	}
}

func (rd *keyReader) PressAnyKey() {
	rd.poll()
}
