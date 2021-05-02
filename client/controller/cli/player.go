package cli

import (
	"github.com/nsf/termbox-go"
	"github.com/simp7/nonograminGo/nonogram"
)

type player struct {
	problemPosition Pos
	position        Pos
	playerMap       [][]Signal
	bitmap          [][]bool
	color           Color
}

/*
	This function Creates player structure with problem position, Width, and Height of map.
	This function will be called when player enter the game or create the map.
*/

func Player(config Color, problemPosition Pos, width int, height int) *player {

	p := new(player)
	p.problemPosition = problemPosition

	p.initMap(width, height)

	p.position = p.problemPosition.Move(0, 1)
	p.color = config

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

func (p *player) SetCell(s Signal) {

	setCell := func(first rune, second rune, fg termbox.Attribute, bg termbox.Attribute) {
		termbox.SetCell(p.position.X, p.position.Y, first, fg, bg)
		termbox.SetCell(p.position.X+1, p.position.Y, second, fg, bg)
	}

	switch s {
	case Empty:
		setCell(' ', ' ', p.color.Empty, p.color.Empty)

	case Fill:
		setCell(' ', ' ', p.color.Filled, p.color.Filled)

	case Check:
		setCell('>', '<', p.color.Checked, p.color.Empty)

	case Wrong:
		setCell('>', '<', p.color.Wrong, p.color.Empty)

	case Cursor:
		setCell('(', ')', p.color.Filled, p.color.Empty)

	case CursorFilled:
		setCell('(', ')', p.color.Empty, p.color.Filled)

	case CursorChecked:
		setCell('(', ')', p.color.Checked, p.color.Empty)

	case CursorWrong:
		setCell('(', ')', p.color.Wrong, p.color.Empty)
	}
}

/*
	This function especially set cursor among cells.
	This function will be called when player move cursor.
*/

func (p *player) SetCursor(cellState Signal) {
	switch cellState {
	case Fill:
		p.SetCell(CursorFilled)
	case Check:
		p.SetCell(CursorChecked)
	case Wrong:
		p.SetCell(CursorWrong)
	default:
		p.SetCell(Cursor)
	}
}

//This function returns real position of the map by calculating cursor position and problem position.

func (p *player) RealPos() (realPos Pos) {
	tmp := p.position.Move(-p.problemPosition.X, -p.problemPosition.Y-1)
	return Pos{tmp.X / 2, tmp.Y}
}

//This function returns current state of current cell of cursor

func (p *player) GetMapSignal() Signal {
	realPos := p.RealPos()
	return p.playerMap[realPos.Y][realPos.X]
}

//This function change state of cell in map

func (p *player) SetMapSignal(signal Signal) {
	realPos := p.RealPos()
	p.playerMap[realPos.Y][realPos.X] = signal
}

// Toggle is called when state of selected cell changed

func (p *player) Toggle(s Signal) {
	p.SetMapSignal(s)
	switch s {
	case Fill:
		p.SetCell(CursorFilled)
	case Check:
		p.SetCell(CursorChecked)
	case Wrong:
		p.SetCell(CursorWrong)
	case Empty:
		p.SetCell(Cursor)
	}
}

/*
	This function process movement of cursor with the help of SetCursor
	This function will be called when cursor moves
*/

func (p *player) moveCursor(condition bool, x int, y int) {
	if condition {
		p.SetCell(p.GetMapSignal())
		p.position = p.position.Move(x, y)
		p.SetCursor(p.GetMapSignal())
	}
}

/*
	This function process movement of cursor with the help of moveCursor
	This function will be called when cursor moves
*/

func (p *player) Move(d Direction) {
	switch d {
	case Up:
		p.moveCursor(p.position.Y-1 >= p.problemPosition.Y+1, 0, -1)
	case Down:
		p.moveCursor(p.position.Y+1 < p.problemPosition.Y+1+len(p.playerMap), 0, 1)
	case Left:
		p.moveCursor(p.position.X-2 >= p.problemPosition.X, -2, 0)
	case Right:
		p.moveCursor(p.position.X+2 < p.problemPosition.X+(2*len(p.playerMap[0])), 2, 0)
	}
}

/*
	This function process playerMap into Bitmap that is just composed with Fill and empty.
	This function will be called when user finish making map in create mode.
*/

func (p *player) FinishCreating(prototype nonogram.Map) nonogram.Map {

	p.bitmap = make([][]bool, len(p.playerMap))
	for n := range p.bitmap {
		p.bitmap[n] = make([]bool, len(p.playerMap[0]))
		p.convertByRow(n)
	}

	return prototype.CopyWithBitmap(p.bitmap)

}

func (p *player) convertByRow(y int) {
	for x := range p.bitmap[y] {
		p.bitmap[y][x] = p.playerMap[y][x] == Fill
	}
}
