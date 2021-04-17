package errs

import (
	"github.com/nsf/termbox-go"
	"io"
	"log"
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
