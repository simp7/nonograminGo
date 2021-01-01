package util

import (
	"errors"
	"github.com/nsf/termbox-go"
	"io"
	"log"
)

var (
	InvalidMap = errors.New("map file has been broken")
)

func CheckErr(e error) {
	if e != nil && e != io.EOF {
		if termbox.IsInit {
			termbox.Close()
		}
		log.Fatal(e)
	}
}
