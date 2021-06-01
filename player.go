package main

import (
	"github.com/nsf/termbox-go"
	"github.com/simp7/nonogram"
	"github.com/simp7/nonogram/unit"
)

type player struct {
	problemPosition Pos
	position        Pos
	playerMap       [][]signal
	bitmap          [][]bool
	color           Color
	core            nonogram.Core
}

//Player returns in-play logic of nonogram.
func Player(config Color, problemPosition Pos, width int, height int, core nonogram.Core) *player {

	p := new(player)
	p.problemPosition = problemPosition

	p.initMap(width, height)

	p.position = p.problemPosition.Move(0, 1)
	p.color = config
	p.core = core

	return p
}

func (p *player) initMap(width int, height int) {
	p.playerMap = make([][]signal, height)
	for n := range p.playerMap {
		p.playerMap[n] = make([]signal, width)
		for m := range p.playerMap[n] {
			p.playerMap[n][m] = Empty
		}
	}
}

func (p *player) SetCell(s signal) {

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

func (p *player) SetCursor(cellState signal) {
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

func (p *player) RealPos() (realPos Pos) {
	tmp := p.position.Move(-p.problemPosition.X, -p.problemPosition.Y-1)
	return Pos{tmp.X / 2, tmp.Y}
}

func (p *player) GetMapSignal() signal {
	realPos := p.RealPos()
	return p.playerMap[realPos.Y][realPos.X]
}

func (p *player) SetMapSignal(signal signal) {
	realPos := p.RealPos()
	p.playerMap[realPos.Y][realPos.X] = signal
}

func (p *player) Toggle(s signal) {
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

func (p *player) moveCursor(condition bool, x int, y int) {
	if condition {
		p.SetCell(p.GetMapSignal())
		p.position = p.position.Move(x, y)
		p.SetCursor(p.GetMapSignal())
	}
}

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

func (p *player) FinishCreating() unit.Map {

	p.bitmap = make([][]bool, len(p.playerMap))
	for n := range p.bitmap {
		p.bitmap[n] = make([]bool, len(p.playerMap[0]))
		p.convertByRow(n)
	}

	return p.core.InitMap(p.bitmap)

}

func (p *player) convertByRow(y int) {
	for x := range p.bitmap[y] {
		p.bitmap[y][x] = p.playerMap[y][x] == Fill
	}
}
