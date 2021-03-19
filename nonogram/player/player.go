package player

import (
	"github.com/nsf/termbox-go"
	"github.com/simp7/nonograminGo/nonogram"
	"github.com/simp7/nonograminGo/nonogram/direction"
	"github.com/simp7/nonograminGo/nonogram/setting"
	"github.com/simp7/nonograminGo/nonogram/signal"
)

type player struct {
	xProblemPos int
	yProblemPos int
	xPos        int
	yPos        int
	playerMap   [][]nonogram.Signal
	bitmap      [][]bool
	*setting.Setting
}

/*
	This function Creates player structure with problem position, Width, and Height of map.
	This function will be called when player enter the game or create the map.
*/

func New(x int, y int, width int, height int) nonogram.Player {

	p := new(player)
	p.xProblemPos, p.yProblemPos = x, y

	p.initMap(width, height)

	p.xPos, p.yPos = p.xProblemPos, p.yProblemPos+1
	p.Setting = setting.Get()

	return p
}

func (p *player) initMap(width int, height int) {
	p.playerMap = make([][]nonogram.Signal, height)
	for n := range p.playerMap {
		p.playerMap[n] = make([]nonogram.Signal, width)
		for m := range p.playerMap[n] {
			p.playerMap[n][m] = signal.Empty
		}
	}
}

/*
	This function set a cell that cursor exist.
	This function will be called when player inputs key in game or in create mode
*/

func (p *player) SetMap(s nonogram.Signal) {

	setCell := func(first rune, second rune, fg termbox.Attribute, bg termbox.Attribute) {
		termbox.SetCell(p.xPos, p.yPos, first, fg, bg)
		termbox.SetCell(p.xPos+1, p.yPos, second, fg, bg)
	}

	switch s {
	case signal.Empty:
		setCell(' ', ' ', p.Empty, p.Empty)

	case signal.Fill:
		setCell(' ', ' ', p.Filled, p.Filled)

	case signal.Check:
		setCell('>', '<', p.Checked, p.Empty)

	case signal.Wrong:
		setCell('>', '<', p.Wrong, p.Empty)

	case signal.Cursor:
		setCell('(', ')', p.Filled, p.Empty)

	case signal.CursorFilled:
		setCell('(', ')', p.Empty, p.Filled)

	case signal.CursorChecked:
		setCell('(', ')', p.Checked, p.Empty)

	case signal.CursorWrong:
		setCell('(', ')', p.Wrong, p.Empty)
	}
}

/*
	This function especially set cursor among cells.
	This function will be called when player move cursor.
*/

func (p *player) SetCursor(cellState nonogram.Signal) {
	switch cellState {
	case signal.Fill:
		p.SetMap(signal.CursorFilled)
	case signal.Check:
		p.SetMap(signal.CursorChecked)
	case signal.Wrong:
		p.SetMap(signal.CursorWrong)
	default:
		p.SetMap(signal.Cursor)
	}
}

//This function returns real position of the map by calculating cursor position and problem position.

func (p *player) RealPos() (realXPos int, realYPos int) {
	realXPos, realYPos = (p.xPos-p.xProblemPos)/2, p.yPos-p.yProblemPos-1
	return
}

//This function returns current state of current cell of cursor

func (p *player) GetMapSignal() nonogram.Signal {
	realXPos, realYPos := p.RealPos()
	return p.playerMap[realYPos][realXPos]
}

//This function change state of cell in map

func (p *player) SetMapSignal(signal nonogram.Signal) {
	realXPos, realYPos := p.RealPos()
	p.playerMap[realYPos][realXPos] = signal
}

// Toggle is called when state of selected cell changed

func (p *player) Toggle(s nonogram.Signal) {
	p.SetMapSignal(s)
	switch s {
	case signal.Fill:
		p.SetMap(signal.CursorFilled)
	case signal.Check:
		p.SetMap(signal.CursorChecked)
	case signal.Wrong:
		p.SetMap(signal.CursorWrong)
	case signal.Empty:
		p.SetMap(signal.Cursor)
	}
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

func (p *player) Move(d nonogram.Direction) {
	switch d {
	case direction.Up:
		p.moveCursor(p.yPos-1 >= p.yProblemPos+1, func() { p.yPos-- })
	case direction.Down:
		p.moveCursor(p.yPos+1 < p.yProblemPos+1+len(p.playerMap), func() { p.yPos++ })
	case direction.Left:
		p.moveCursor(p.xPos-2 >= p.xProblemPos, func() { p.xPos -= 2 })
	case direction.Right:
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
		p.bitmap[y][x] = p.playerMap[y][x] == signal.Fill
	}
}