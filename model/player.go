package model

import (
	"../asset"
	"github.com/nsf/termbox-go"
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
	xpos        int
	ypos        int
	playermap   [][]Signal
}

func NewPlayer(x int, y int, width int, height int) *Player {

	pl := Player{}
	pl.xProblemPos, pl.yProblemPos = x, y

	pl.playermap = make([][]Signal, height)
	for n := range pl.playermap {
		pl.playermap[n] = make([]Signal, width)
		for m := range pl.playermap[n] {
			pl.playermap[n][m] = Empty
		}
	}

	pl.xpos, pl.ypos = pl.xProblemPos, pl.yProblemPos+1

	return &pl
}

func (pl *Player) SetMap(signal Signal) {

	setCell := func(first rune, second rune, fg termbox.Attribute, bg termbox.Attribute) {
		termbox.SetCell(pl.xpos, pl.ypos, first, fg, bg)
		termbox.SetCell(pl.xpos+1, pl.ypos, second, fg, bg)
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

func (pl *Player) GetRealpos() (realxpos int, realypos int) {
	realxpos, realypos = (pl.xpos-pl.xProblemPos)/2, pl.ypos-pl.yProblemPos-1
	return
}

func (pl *Player) GetMapSignal() Signal {
	realxpos, realypos := pl.GetRealpos()
	return pl.playermap[realypos][realxpos]
}

func (pl *Player) SetMapSignal(signal Signal) {
	realxpos, realypos := pl.GetRealpos()
	pl.playermap[realypos][realxpos] = signal
}

func (pl *Player) moveCursor(condition bool, function func()) {
	if condition {
		pl.SetMap(pl.GetMapSignal())
		function()
		pl.SetCursor(pl.GetMapSignal())
	}
}

func (pl *Player) Move(direction Direction) {
	switch direction {
	case Up:
		pl.moveCursor(pl.ypos-1 >= pl.yProblemPos+1, func() { pl.ypos-- })
	case Down:
		pl.moveCursor(pl.ypos+1 < pl.yProblemPos+1+len(pl.playermap), func() { pl.ypos++ })
	case Left:
		pl.moveCursor(pl.xpos-2 >= pl.xProblemPos, func() { pl.xpos -= 2 })
	case Right:
		pl.moveCursor(pl.xpos+2 < pl.xProblemPos+(2*len(pl.playermap[0])), func() { pl.xpos += 2 })
	}
}

func (pl *Player) ConvertToBitMap() (result [][]bool) {
	result = make([][]bool, len(pl.playermap))
	for n := range result {
		result[n] = make([]bool, len(pl.playermap[0]))
		for m := range result[n] {
			if pl.playermap[n][m] == Fill {
				result[n][m] = true
			} else {
				result[n][m] = false
			}
		}
	}
	return
}
