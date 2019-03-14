package asset

import "github.com/nsf/termbox-go"

const (
	ColorFilledCell  = termbox.ColorBlack
	ColorEmptyCell   = termbox.ColorWhite
	ColorCheckedCell = termbox.ColorRed | termbox.AttrBold
	ColorText        = termbox.ColorBlack
)

var (
	StringMainMenu = []string{"----------", " NONOGRAM", "----------", "Press number you want to select.", "1. START", "2. LOAD", "3. Create", "4. Exit"}
	StringResult   = []string{"----------", "  RESULT", "----------", "Map name : ", "Map size : ", "Cleared time : ", "Press any key to Continue."}
)
