package model

import (
	"github.com/nsf/termbox-go"
	"github.com/simp7/nonograminGo/asset"
)

type Signal uint8
type Direction uint8

const (
	Empty Signal = iota
	Fill
	Check
	Wrong
	Cursor
	CursorFilled
	CursorChecked
	CursorWrong
)

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Player struct {
	xProblemPos int
	yProblemPos int
	xPos        int
	yPos        int
	playerMap   [][]Signal
}

/*
	This function Creates player structure with problem position, width, and height of map.
	This function will be called when player enter the game or create the map.
*/

func NewPlayer(x int, y int, width int, height int) *Player {

	pl := Player{}
	pl.xProblemPos, pl.yProblemPos = x, y

	pl.playerMap = make([][]Signal, height)
	for n := range pl.playerMap {
		pl.playerMap[n] = make([]Signal, width)
		for m := range pl.playerMap[n] {
			pl.playerMap[n][m] = Empty
		}
	}

	pl.xPos, pl.yPos = pl.xProblemPos, pl.yProblemPos+1

	return &pl
}

/*
	This function set a cell that cursor exist.
	This function will be called when player inputs key in game or in create mode
*/

func (pl *Player) SetMap(signal Signal) {

	setCell := func(first rune, second rune, fg termbox.Attribute, bg termbox.Attribute) {
		termbox.SetCell(pl.xPos, pl.yPos, first, fg, bg)
		termbox.SetCell(pl.xPos+1, pl.yPos, second, fg, bg)
	}

	switch signal {
	case Empty:
		setCell(' ', ' ', asset.ColorEmptyCell, asset.ColorEmptyCell)

	case Fill:
		setCell(' ', ' ', asset.ColorFilledCell, asset.ColorFilledCell)

	case Check:
		setCell('>', '<', asset.ColorCheckedCell, asset.ColorEmptyCell)

	case Wrong:
		setCell('>', '<', asset.ColorWrongCell, asset.ColorEmptyCell)

	case Cursor:
		setCell('(', ')', asset.ColorFilledCell, asset.ColorEmptyCell)

	case CursorFilled:
		setCell('(', ')', asset.ColorEmptyCell, asset.ColorFilledCell)

	case CursorChecked:
		setCell('(', ')', asset.ColorCheckedCell, asset.ColorEmptyCell)

	case CursorWrong:
		setCell('(', ')', asset.ColorWrongCell, asset.ColorEmptyCell)
	}
}

/*
	This function especially set cursor among cells.
	This function will be called when player move cursor.
*/

func (pl *Player) SetCursor(cellState Signal) {
	if cellState == Fill {
		pl.SetMap(CursorFilled)
	} else if cellState == Check {
		pl.SetMap(CursorChecked)
	} else if cellState == Wrong {
		pl.SetMap(CursorWrong)
	} else {
		pl.SetMap(Cursor)
	}
}

//This function returns real position of the map by calculating cursor position and problem position.

func (pl *Player) RealPos() (realXPos int, realYPos int) {
	realXPos, realYPos = (pl.xPos-pl.xProblemPos)/2, pl.yPos-pl.yProblemPos-1
	return
}

//This function returns current state of current cell of cursor

func (pl *Player) GetMapSignal() Signal {
	realXPos, realYPos := pl.RealPos()
	return pl.playerMap[realYPos][realXPos]
}

//This function change state of cell in map

func (pl *Player) SetMapSignal(signal Signal) {
	realXPos, realYPos := pl.RealPos()
	pl.playerMap[realYPos][realXPos] = signal
}

/*
	This function process movement of cursor with the help of SetCursor
	This function will be called when cursor moves
*/

func (pl *Player) moveCursor(condition bool, function func()) {
	if condition {
		pl.SetMap(pl.GetMapSignal())
		function()
		pl.SetCursor(pl.GetMapSignal())
	}
}

/*
	This function process movement of cursor with the help of moveCursor
	This function will be called when cursor moves
*/

func (pl *Player) Move(direction Direction) {
	switch direction {
	case Up:
		pl.moveCursor(pl.yPos-1 >= pl.yProblemPos+1, func() { pl.yPos-- })
	case Down:
		pl.moveCursor(pl.yPos+1 < pl.yProblemPos+1+len(pl.playerMap), func() { pl.yPos++ })
	case Left:
		pl.moveCursor(pl.xPos-2 >= pl.xProblemPos, func() { pl.xPos -= 2 })
	case Right:
		pl.moveCursor(pl.xPos+2 < pl.xProblemPos+(2*len(pl.playerMap[0])), func() { pl.xPos += 2 })
	}
}

/*
	This function process playerMap into bitmap that is just composed with Fill and empty.
	This function will be called when user finish making map in create mode.
*/

func (pl *Player) ConvertToBitMap() (result [][]bool) {
	result = make([][]bool, len(pl.playerMap))
	for n := range result {
		result[n] = make([]bool, len(pl.playerMap[0]))
		for m := range result[n] {
			if pl.playerMap[n][m] == Fill {
				result[n][m] = true
			} else {
				result[n][m] = false
			}
		}
	}
	return
}
