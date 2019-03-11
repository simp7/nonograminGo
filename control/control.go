package control

/*
import (
	"bufio"
	"os"
)

// This file deals with user's control.

type keyStroker struct {
	reader *bufio.Reader
	key    byte
	err    error
}

func NewKeyStroker() *keyStroker {
	rd := keyStroker{}
	rd.reader = bufio.NewReader(os.Stdin)
	rd.key = 0
	rd.err = nil
	return &rd
}

func (rd *keyStroker) ControlMenu() {
	for {
		rd.key, rd.err = rd.reader.ReadByte()
		CheckErr(rd.err)
		switch rd.key {
		case '1':
			ShowMapList()
			rd.SelectMap()
		case '2':

		case '3':
			ShowResult()
			rd.PressAnyKey()
		case '4':
			return
		}
	}
}

func (rd *keyStroker) SelectMap() {
	var mapName []byte
	for {
		mapName, _, rd.err = rd.reader.ReadLine()
		CheckErr(rd.err)
		targetMap := newNonomap(string(mapName))
	}
}

func (rd *keyStroker) ControlGame() {
	for {
		rd.key, rd.err = rd.reader.ReadByte()
		CheckErr(rd.err)

	}
}

func (rd *keyStroker) PressAnyKey() { //calls when entering hall of honor
	rd.reader.ReadByte()
}
*/
