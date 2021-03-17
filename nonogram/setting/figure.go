package setting

import (
	"github.com/simp7/nonograminGo/nonogram/position"
)

type Figure struct {
	NameMax    int
	WidthMax   int
	HeightMax  int
	DefaultX   int
	DefaultY   int
	DefaultPos position.Pos
}
