package asset

import (
	"github.com/nsf/termbox-go"
)

const (
	ColorFilledCell  = termbox.ColorBlack
	ColorEmptyCell   = termbox.ColorWhite
	ColorCheckedCell = termbox.ColorCyan
	ColorWrongCell   = termbox.ColorRed | termbox.AttrBold
	ColorText        = termbox.ColorBlack
)

var (
	StringMainMenu     = []string{"----------", " NONOGRAM", "----------", "", "Press number you want to select.", "", "1. START", "2. CREATE", "3. CREDIT", "4. EXIT"}
	StringSelectHeader = []string{"[mapList]  [<-Prev | Next->]", "----------------------------", ""}
	StringResult       = []string{"--------------------", "       CLEAR!", "--------------------", "", "MAP NAME    : ", "CLEAR TIME  : ", "WRONG CELLS : "}
	StringCredit       = []string{"--------------------------------------", "                CREDIT", "--------------------------------------", "Developer : JeongHyeon Park(N0RM4L15T)", "Programming Language : Go(100%)", "I wish you enjoy this game!", "--------------------------------------"}
)

var (
	StringHeaderMapname   = "Write map name that you want to create"
	StringHeaderWidth     = "Write map's width"
	StringHeaderHeight    = "Write map's height"
	StringMsgFileNotExist = "File doesn't exist."
)

var (
	NumberNameMax  = 30
	NumberDefaultX = 5
	NumberDefaultY = 5
)
