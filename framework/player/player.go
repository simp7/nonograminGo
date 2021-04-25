package player

import (
	"github.com/nsf/termbox-go"
	"github.com/simp7/nonograminGo/framework"
	direction2 "github.com/simp7/nonograminGo/framework/direction"
	signal2 "github.com/simp7/nonograminGo/framework/signal"
	"github.com/simp7/nonograminGo/nonogram"
	"github.com/simp7/nonograminGo/nonogram/standard"
)

type player struct {
	xProblemPos int
	yProblemPos int
	xPos        int
	yPos        int
	playerMap   [][]framework.Signal
	bitmap      [][]bool
	color       framework.Color
}

/*
	This function Creates player structure with problem position, Width, and Height of map.
	This function will be called when player enter the game or create the map.
*/

func New(config framework.Color, x int, y int, width int, height int) framework.Player {

	p := new(player)
	p.xProblemPos, p.yProblemPos = x, y

	p.initMap(width, height)

	p.xPos, p.yPos = p.xProblemPos, p.yProblemPos+1
	p.color = config

	return p
}

func (p *player) initMap(width int, height int) {
	p.playerMap = make([][]framework.Signal, height)
	for n := range p.playerMap {
		p.playerMap[n] = make([]framework.Signal, width)
		for m := range p.playerMap[n] {
			p.playerMap[n][m] = signal2.Empty
		}
	}
}

/*
	This function set a cell that cursor exist.
	This function will be called when player inputs key in game or in create mode
*/

func (p *player) SetCell(s framework.Signal) {

	setCell := func(first rune, second rune, fg termbox.Attribute, bg termbox.Attribute) {
		termbox.SetCell(p.xPos, p.yPos, first, fg, bg)
		termbox.SetCell(p.xPos+1, p.yPos, second, fg, bg)
	}

	switch s {
	case signal2.Empty:
		setCell(' ', ' ', p.color.Empty, p.color.Empty)

	case signal2.Fill:
		setCell(' ', ' ', p.color.Filled, p.color.Filled)

	case signal2.Check:
		setCell('>', '<', p.color.Checked, p.color.Empty)

	case signal2.Wrong:
		setCell('>', '<', p.color.Wrong, p.color.Empty)

	case signal2.Cursor:
		setCell('(', ')', p.color.Filled, p.color.Empty)

	case signal2.CursorFilled:
		setCell('(', ')', p.color.Empty, p.color.Filled)

	case signal2.CursorChecked:
		setCell('(', ')', p.color.Checked, p.color.Empty)

	case signal2.CursorWrong:
		setCell('(', ')', p.color.Wrong, p.color.Empty)
	}
}

/*
	This function especially set cursor among cells.
	This function will be called when player move cursor.
*/

func (p *player) SetCursor(cellState framework.Signal) {
	switch cellState {
	case signal2.Fill:
		p.SetCell(signal2.CursorFilled)
	case signal2.Check:
		p.SetCell(signal2.CursorChecked)
	case signal2.Wrong:
		p.SetCell(signal2.CursorWrong)
	default:
		p.SetCell(signal2.Cursor)
	}
}

//This function returns real position of the map by calculating cursor position and problem position.

func (p *player) RealPos() (realXPos int, realYPos int) {
	realXPos, realYPos = (p.xPos-p.xProblemPos)/2, p.yPos-p.yProblemPos-1
	return
}

//This function returns current state of current cell of cursor

func (p *player) GetMapSignal() framework.Signal {
	realXPos, realYPos := p.RealPos()
	return p.playerMap[realYPos][realXPos]
}

//This function change state of cell in map

func (p *player) SetMapSignal(signal framework.Signal) {
	realXPos, realYPos := p.RealPos()
	p.playerMap[realYPos][realXPos] = signal
}

// Toggle is called when state of selected cell changed

func (p *player) Toggle(s framework.Signal) {
	p.SetMapSignal(s)
	switch s {
	case signal2.Fill:
		p.SetCell(signal2.CursorFilled)
	case signal2.Check:
		p.SetCell(signal2.CursorChecked)
	case signal2.Wrong:
		p.SetCell(signal2.CursorWrong)
	case signal2.Empty:
		p.SetCell(signal2.Cursor)
	}
}

/*
	This function process movement of cursor with the help of SetCursor
	This function will be called when cursor moves
*/

func (p *player) moveCursor(condition bool, function func()) {
	if condition {
		p.SetCell(p.GetMapSignal())
		function()
		p.SetCursor(p.GetMapSignal())
	}
}

/*
	This function process movement of cursor with the help of moveCursor
	This function will be called when cursor moves
*/

func (p *player) Move(d framework.Direction) {
	switch d {
	case direction2.Up:
		p.moveCursor(p.yPos-1 >= p.yProblemPos+1, func() { p.yPos-- })
	case direction2.Down:
		p.moveCursor(p.yPos+1 < p.yProblemPos+1+len(p.playerMap), func() { p.yPos++ })
	case direction2.Left:
		p.moveCursor(p.xPos-2 >= p.xProblemPos, func() { p.xPos -= 2 })
	case direction2.Right:
		p.moveCursor(p.xPos+2 < p.xProblemPos+(2*len(p.playerMap[0])), func() { p.xPos += 2 })
	}
}

/*
	This function process playerMap into Bitmap that is just composed with Fill and empty.
	This function will be called when user finish making map in create mode.
*/

func (p *player) FinishCreating() nonogram.Map {

	p.bitmap = make([][]bool, len(p.playerMap))
	for n := range p.bitmap {
		p.bitmap[n] = make([]bool, len(p.playerMap[0]))
		p.convertByRow(n)
	}

	return standard.NewByBitMap(p.bitmap)

}

func (p *player) convertByRow(y int) {
	for x := range p.bitmap[y] {
		p.bitmap[y][x] = p.playerMap[y][x] == signal2.Fill
	}
}
