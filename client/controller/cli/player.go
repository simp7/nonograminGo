package cli

import (
	"github.com/nsf/termbox-go"
	"github.com/simp7/nonograminGo/client"
	"github.com/simp7/nonograminGo/nonogram"
	"github.com/simp7/nonograminGo/nonogram/standard"
)

type player struct {
	xProblemPos int
	yProblemPos int
	xPos        int
	yPos        int
	playerMap   [][]client.Signal
	bitmap      [][]bool
	color       client.Color
}

/*
	This function Creates player structure with problem position, Width, and Height of map.
	This function will be called when player enter the game or create the map.
*/

func Player(config client.Color, x int, y int, width int, height int) client.Player {

	p := new(player)
	p.xProblemPos, p.yProblemPos = x, y

	p.initMap(width, height)

	p.xPos, p.yPos = p.xProblemPos, p.yProblemPos+1
	p.color = config

	return p
}

func (p *player) initMap(width int, height int) {
	p.playerMap = make([][]client.Signal, height)
	for n := range p.playerMap {
		p.playerMap[n] = make([]client.Signal, width)
		for m := range p.playerMap[n] {
			p.playerMap[n][m] = client.Empty
		}
	}
}

/*
	This function set a cell that cursor exist.
	This function will be called when player inputs key in game or in create mode
*/

func (p *player) SetCell(s client.Signal) {

	setCell := func(first rune, second rune, fg termbox.Attribute, bg termbox.Attribute) {
		termbox.SetCell(p.xPos, p.yPos, first, fg, bg)
		termbox.SetCell(p.xPos+1, p.yPos, second, fg, bg)
	}

	switch s {
	case client.Empty:
		setCell(' ', ' ', p.color.Empty, p.color.Empty)

	case client.Fill:
		setCell(' ', ' ', p.color.Filled, p.color.Filled)

	case client.Check:
		setCell('>', '<', p.color.Checked, p.color.Empty)

	case client.Wrong:
		setCell('>', '<', p.color.Wrong, p.color.Empty)

	case client.Cursor:
		setCell('(', ')', p.color.Filled, p.color.Empty)

	case client.CursorFilled:
		setCell('(', ')', p.color.Empty, p.color.Filled)

	case client.CursorChecked:
		setCell('(', ')', p.color.Checked, p.color.Empty)

	case client.CursorWrong:
		setCell('(', ')', p.color.Wrong, p.color.Empty)
	}
}

/*
	This function especially set cursor among cells.
	This function will be called when player move cursor.
*/

func (p *player) SetCursor(cellState client.Signal) {
	switch cellState {
	case client.Fill:
		p.SetCell(client.CursorFilled)
	case client.Check:
		p.SetCell(client.CursorChecked)
	case client.Wrong:
		p.SetCell(client.CursorWrong)
	default:
		p.SetCell(client.Cursor)
	}
}

//This function returns real position of the map by calculating cursor position and problem position.

func (p *player) RealPos() (realXPos int, realYPos int) {
	realXPos, realYPos = (p.xPos-p.xProblemPos)/2, p.yPos-p.yProblemPos-1
	return
}

//This function returns current state of current cell of cursor

func (p *player) GetMapSignal() client.Signal {
	realXPos, realYPos := p.RealPos()
	return p.playerMap[realYPos][realXPos]
}

//This function change state of cell in map

func (p *player) SetMapSignal(signal client.Signal) {
	realXPos, realYPos := p.RealPos()
	p.playerMap[realYPos][realXPos] = signal
}

// Toggle is called when state of selected cell changed

func (p *player) Toggle(s client.Signal) {
	p.SetMapSignal(s)
	switch s {
	case client.Fill:
		p.SetCell(client.CursorFilled)
	case client.Check:
		p.SetCell(client.CursorChecked)
	case client.Wrong:
		p.SetCell(client.CursorWrong)
	case client.Empty:
		p.SetCell(client.Cursor)
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

func (p *player) Move(d client.Direction) {
	switch d {
	case client.Up:
		p.moveCursor(p.yPos-1 >= p.yProblemPos+1, func() { p.yPos-- })
	case client.Down:
		p.moveCursor(p.yPos+1 < p.yProblemPos+1+len(p.playerMap), func() { p.yPos++ })
	case client.Left:
		p.moveCursor(p.xPos-2 >= p.xProblemPos, func() { p.xPos -= 2 })
	case client.Right:
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
		p.bitmap[y][x] = p.playerMap[y][x] == client.Fill
	}
}
