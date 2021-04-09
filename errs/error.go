package errs

import (
	"errors"
	"github.com/nsf/termbox-go"
	"io"
	"log"
)

var (
	InvalidMap  = errors.New("map file has been broken")
	InvalidType = errors.New("type of two didn't match")
	TooManyArgs = errors.New("argument should be less than 2")
)

func Check(e error) {

	if e == nil || e == io.EOF {
		return
	}

	if termbox.IsInit {
		termbox.Close()
	}
	log.Fatal(e)

}
