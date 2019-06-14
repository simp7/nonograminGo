package asset

import (
	"github.com/nsf/termbox-go"
)

const (
	ColorFilledCell  = termbox.ColorBlack
	ColorEmptyCell   = termbox.ColorWhite
	ColorCheckedCell = termbox.ColorRed | termbox.AttrBold
	ColorText        = termbox.ColorBlack
)

var (
	StringMainMenu     = []string{"----------", " NONOGRAM", "----------", "", "Press number you want to select.", "", "1. START", "2. CREATE", "3. CREDIT", "4. EXIT"}
	StringSelectHeader = []string{"[mapList]  [<-Prev | Next->]"}
	StringResult       = []string{"--------------------", "       CLEAR!", "--------------------", "", "MAP NAME    : ", "CLEAR TIME  : ", "WRONG CELLS : "}
)
