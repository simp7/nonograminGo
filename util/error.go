package util

import (
	"errors"
	"github.com/nsf/termbox-go"
	"log"
)

var (
	InvalidMap = errors.New("Map file has been broken.")
)

func CheckErr(e error) {
	if e != nil {
		termbox.Close()
		log.Fatal(e)
	}
}
