package player

import (
	"github.com/nsf/termbox-go"
	"github.com/simp7/nonograminGo/nonogram"
	"github.com/simp7/nonograminGo/nonogram/setting"
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
			p.playerMap[n][m] = nonogram.Empty
		}
	}
}

/*
	This function set a cell that cursor exist.
	This function will be called when player inputs key in game or in create mode
*/

func (p *player) SetMap(signal nonogram.Signal) {

	setCell := func(first rune, second rune, fg termbox.Attribute, bg termbox.Attribute) {
		termbox.SetCell(p.xPos, p.yPos, first, fg, bg)
		termbox.SetCell(p.xPos+1, p.yPos, second, fg, bg)
	}

	switch signal {
	case nonogram.Empty:
		setCell(' ', ' ', p.Empty, p.Empty)

	case nonogram.Fill:
		setCell(' ', ' ', p.Filled, p.Filled)

	case nonogram.Check:
		setCell('>', '<', p.Checked, p.Empty)

	case nonogram.Wrong:
		setCell('>', '<', p.Wrong, p.Empty)

	case nonogram.Cursor:
		setCell('(', ')', p.Filled, p.Empty)

	case nonogram.CursorFilled:
		setCell('(', ')', p.Empty, p.Filled)

	case nonogram.CursorChecked:
		setCell('(', ')', p.Checked, p.Empty)

	case nonogram.CursorWrong:
		setCell('(', ')', p.Wrong, p.Empty)
	}
}

/*
	This function especially set cursor among cells.
	This function will be called when player move cursor.
*/

func (p *player) SetCursor(cellState nonogram.Signal) {
	switch cellState {
	case nonogram.Fill:
		p.SetMap(nonogram.CursorFilled)
	case nonogram.Check:
		p.SetMap(nonogram.CursorChecked)
	case nonogram.Wrong:
		p.SetMap(nonogram.CursorWrong)
	default:
		p.SetMap(nonogram.Cursor)
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

func (p *player) Toggle(signal nonogram.Signal) {
	p.SetMapSignal(signal)
	switch signal {
	case nonogram.Fill:
		p.SetMap(nonogram.CursorFilled)
	case nonogram.Check:
		p.SetMap(nonogram.CursorChecked)
	case nonogram.Wrong:
		p.SetMap(nonogram.CursorWrong)
	case nonogram.Empty:
		p.SetMap(nonogram.Cursor)
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

func (p *player) Move(direction nonogram.Direction) {
	switch direction {
	case nonogram.Up:
		p.moveCursor(p.yPos-1 >= p.yProblemPos+1, func() { p.yPos-- })
	case nonogram.Down:
		p.moveCursor(p.yPos+1 < p.yProblemPos+1+len(p.playerMap), func() { p.yPos++ })
	case nonogram.Left:
		p.moveCursor(p.xPos-2 >= p.xProblemPos, func() { p.xPos -= 2 })
	case nonogram.Right:
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
		p.bitmap[y][x] = p.playerMap[y][x] == nonogram.Fill
	}
}
