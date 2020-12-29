package util

import (
	"errors"
	"github.com/nsf/termbox-go"
	"log"
)

var (
	InvalidMap = errors.New("map file has been broken")
)

func CheckErr(e error) {
	if e != nil {
		if termbox.IsInit {
			termbox.Close()
		}
		log.Fatal(e)
	}
}
