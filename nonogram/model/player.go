package model

import (
	"github.com/nsf/termbox-go"
	"github.com/simp7/nonograminGo/nonogram/asset"
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

type Player interface {
	SetMap(Signal)
	SetCursor(Signal)
	RealPos() (x, y int)
	GetMapSignal() Signal
	SetMapSignal(Signal)
	Move(Direction)
	FinishCreating() [][]bool
}

type player struct {
	xProblemPos int
	yProblemPos int
	xPos        int
	yPos        int
	playerMap   [][]Signal
	bitmap      [][]bool
	*asset.Setting
}

/*
	This function Creates player structure with problem position, Width, and Height of map.
	This function will be called when player enter the game or create the map.
*/

func NewPlayer(x int, y int, width int, height int) Player {

	p := new(player)
	p.xProblemPos, p.yProblemPos = x, y

	p.initMap(width, height)

	p.xPos, p.yPos = p.xProblemPos, p.yProblemPos+1
	p.Setting = asset.GetSetting()

	return p
}

func (p *player) initMap(width int, height int) {
	p.playerMap = make([][]Signal, height)
	for n := range p.playerMap {
		p.playerMap[n] = make([]Signal, width)
		for m := range p.playerMap[n] {
			p.playerMap[n][m] = Empty
		}
	}
}

/*
	This function set a cell that cursor exist.
	This function will be called when player inputs key in game or in create mode
*/

func (p *player) SetMap(signal Signal) {

	setCell := func(first rune, second rune, fg termbox.Attribute, bg termbox.Attribute) {
		termbox.SetCell(p.xPos, p.yPos, first, fg, bg)
		termbox.SetCell(p.xPos+1, p.yPos, second, fg, bg)
	}

	switch signal {
	case Empty:
		setCell(' ', ' ', p.Empty, p.Empty)

	case Fill:
		setCell(' ', ' ', p.Filled, p.Filled)

	case Check:
		setCell('>', '<', p.Checked, p.Empty)

	case Wrong:
		setCell('>', '<', p.Wrong, p.Empty)

	case Cursor:
		setCell('(', ')', p.Filled, p.Empty)

	case CursorFilled:
		setCell('(', ')', p.Empty, p.Filled)

	case CursorChecked:
		setCell('(', ')', p.Checked, p.Empty)

	case CursorWrong:
		setCell('(', ')', p.Wrong, p.Empty)
	}
}

/*
	This function especially set cursor among cells.
	This function will be called when player move cursor.
*/

func (p *player) SetCursor(cellState Signal) {
	switch cellState {
	case Fill:
		p.SetMap(CursorFilled)
	case Check:
		p.SetMap(Check)
	case Wrong:
		p.SetMap(CursorWrong)
	default:
		p.SetMap(Cursor)
	}
}

//This function returns real position of the map by calculating cursor position and problem position.

func (p *player) RealPos() (realXPos int, realYPos int) {
	realXPos, realYPos = (p.xPos-p.xProblemPos)/2, p.yPos-p.yProblemPos-1
	return
}

//This function returns current state of current cell of cursor

func (p *player) GetMapSignal() Signal {
	realXPos, realYPos := p.RealPos()
	return p.playerMap[realYPos][realXPos]
}

//This function change state of cell in map

func (p *player) SetMapSignal(signal Signal) {
	realXPos, realYPos := p.RealPos()
	p.playerMap[realYPos][realXPos] = signal
}

/*
	This function process movement of cursor with the help of SetCursor
	This function will be called when cursor moves
*/

func (p *player) moveCursor(condition bool, function func()) {
	if condition {
		p.SetMap(p.GetMapSignal())
		function()
		p.SetCursor(p.GetMapSignal())
	}
}

/*
	This function process movement of cursor with the help of moveCursor
	This function will be called when cursor moves
*/

func (p *player) Move(direction Direction) {
	switch direction {
	case Up:
		p.moveCursor(p.yPos-1 >= p.yProblemPos+1, func() { p.yPos-- })
	case Down:
		p.moveCursor(p.yPos+1 < p.yProblemPos+1+len(p.playerMap), func() { p.yPos++ })
	case Left:
		p.moveCursor(p.xPos-2 >= p.xProblemPos, func() { p.xPos -= 2 })
	case Right:
		p.moveCursor(p.xPos+2 < p.xProblemPos+(2*len(p.playerMap[0])), func() { p.xPos += 2 })
	}
}

/*
	This function process playerMap into Bitmap that is just composed with Fill and empty.
	This function will be called when user finish making map in create mode.
*/

func (p *player) FinishCreating() [][]bool {

	p.bitmap = make([][]bool, len(p.playerMap))
	for n := range p.bitmap {
		p.bitmap[n] = make([]bool, len(p.playerMap[0]))
		p.convertByRow(n)
	}

	return p.bitmap

}

func (p *player) convertByRow(y int) {
	for x := range p.bitmap[y] {
		p.bitmap[y][x] = p.playerMap[y][x] == Fill
	}
}
