package main

import (
	"bufio"
	"os"
)

// This file deals with user's control.

type keyStroker struct {
	reader *bufio.Reader
	key    int
}

func NewKeyStroker() *keyStroker {
	rd := keyStroker{}
	rd.reader = bufio.NewReader(os.Stdin)
	rd.key = 0
	return &rd
}

func (rd *keyStroker) ControlMenu() {
	for {
		rd.key, _ = rd.reader.ReadByte()
		switch key {
		case '1':
			ShowMapList()
		case '2':

		case '3':
			ShowResult()
		case '4':
			return
		}
	}
}

func (rd *keyStroker) SelectMap() {
}

func (rd *keyStroker) ControlGame() {
}

func (rd *keyStroker) PressAnyKey() { //calls when entering hall of honor
	rd.reader.ReadByte()
}
