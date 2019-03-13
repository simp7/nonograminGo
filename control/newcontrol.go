package control

import (
	"../model"
	"../util"
	"../view"
	"bufio"
	"github.com/nsf/termbox-go"
	"os"
)

type KeyReader struct {
	eventChan chan termbox.Event
	endChan   chan bool
}

func NewKeyReader() *KeyReader {

	rd := KeyReader{}
	rd.eventChan = make(chan termbox.Event)
	rd.endChan = make(chan bool)
	return &rd

}

/*
This function takes user input into channel.
This function will be called when program starts.
This function should be called as goroutine.
*/

func (rd *KeyReader) Control() {

	err := termbox.Init()
	util.CheckErr(err)
	defer termbox.Close()

	for {
		select {
		case rd.eventChan <- termbox.PollEvent():
		case <-rd.endChan:
			close(rd.eventChan)
			close(rd.endChan)
			return
		}
	}

}

func (rd *KeyReader) ControlMenu() {

	for {

		event := <-rd.eventChan

		if event.Type == termbox.EventKey {

			switch {
			case event.Ch == '1':
				view.ShowMapList()
				rd.SelectMap()
			case event.Ch == '2':
			case event.Ch == '3':
			case event.Ch == '4' || event.Key == termbox.KeyEsc:
				rd.endChan <- true
				return
			}

		}

	}

}

func (rd *KeyReader) SelectMap() {

	nameReader := bufio.NewReader(os.Stdin)

	for {
		mapName, _, err := nameReader.ReadLine()
		util.CheckErr(err)
		targetMap := model.NewNonomap(string(mapName))
		targetMap.CompareMap(*targetMap, 3)
	}

}

func (rd *KeyReader) ControlGame() {
	for {
	}
}

func (rd *KeyReader) PressAnyKey() {
}
