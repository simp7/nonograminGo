package main

import (
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
			showMapList()
			rd.SelectMap()
		case '2':

		case '3':
			showResult()
			rd.PressAnyKey()
		case '4':
			return

		}
	}

}

func (rd *keyReader) SelectMap() {

	nameReader := bufio.NewReader(os.Stdin)

	for {
		mapName, _, err := rd.reader.ReadLine()
		CheckErr(err)
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
