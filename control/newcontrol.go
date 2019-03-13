package control

import (
	"../asset"
	"../model"
	"../util"
	"bufio"
	"github.com/nsf/termbox-go"
	"os"
)

type KeyReader struct {
	eventChan chan termbox.Event
	endChan   chan struct{}
}

func NewKeyReader() *KeyReader {

	rd := KeyReader{}
	rd.eventChan = make(chan termbox.Event)
	rd.endChan = make(chan struct{})
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
	termbox.Clear(asset.Empty, asset.Empty)
}

func (rd *KeyReader) printf(x int, y int, msgs []string) {

	for _, msg := range msgs {
		for _, ch := range msg {
			termbox.SetCell(x, y, ch, asset.Filled, asset.Empty)
			x++
		}
		y++
	}

}

func (rd *KeyReader) menu() {

	for {

		rd.refresh()
		event := <-rd.eventChan

		if event.Type == termbox.EventKey {
			switch {
			case event.Ch == '1':
				rd.selectMap()
			case event.Ch == '2':
			case event.Ch == '3':
			case event.Ch == '4' || event.Key == termbox.KeyEsc:
				close(rd.endChan)
				return
			}

		}

	}

}

func (rd *KeyReader) selectMap() {

	nameReader := bufio.NewReader(os.Stdin)

	for {
		mapName, _, err := nameReader.ReadLine()
		util.CheckErr(err)
		targetMap := model.NewNonomap(string(mapName))
		targetMap.CompareMap(*targetMap, 3)
	}

}

func (rd *KeyReader) controlGame() {
	for {
	}
}
