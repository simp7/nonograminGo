package asset

import "github.com/nsf/termbox-go"

const (
	ColorFilledCell  = termbox.ColorBlack
	ColorEmptyCell   = termbox.ColorWhite
	ColorCheckedCell = termbox.ColorCyan
	ColorWrongCell   = termbox.ColorRed | termbox.AttrBold
	ColorText        = termbox.ColorBlack
)

const (
	NumberNameMax   = 30
	NumberWidthMax  = 30
	NumberHeightMax = 30
	NumberDefaultX  = 5
	NumberDefaultY  = 5
)

var (
	StringMainMenu     = []string{"----------", " NONOGRAM", "----------", "", "Press number you want to select.", "", "1. START", "2. CREATE", "3. HELP", "4. CREDIT", "5. EXIT"}
	StringSelectHeader = []string{"[mapList]  [<-Prev | Next->]    ", "----------------------------", ""}
	StringResult       = []string{"--------------------", "       CLEAR!", "--------------------", "", "MAP NAME    : ", "CLEAR TIME  : ", "WRONG CELLS : "}
	StringCredit       = []string{"--------------------------------------", "                CREDIT", "--------------------------------------", "Developer : JeongHyeon Park(N0RM4L15T)", "Programming Language : Go(100%)", "This program applies the MIT License(MIT)", "I wish you enjoy this game!", "--------------------------------------"}
	StringComplete     = []string{"----------Congratulation! You Complete Me!----------"}
	StringHelp         = []string{"    MANUAL", "--------------", "ArrowKey : Move Cursor", "Space or Z : Fill the cell", "X : Check the cell that is supposed to be blank", "Enter(Create mode) : Save the map that player create", "Esc : get out of current game/display"}
)

const (
	StringHeaderMapname           = "Write map name that you want to create"
	StringHeaderWidth             = "Write map's width"
	StringHeaderHeight            = "Write map's height"
	StringHeaderSizeError         = "Value should not be more than "
	StringMsgFileNotExist         = "File doesn't exist."
	StringBlankBetweenMapAndTimer = "               "
)
