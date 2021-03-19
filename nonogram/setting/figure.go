package setting

import (
	"github.com/simp7/nonograminGo/nonogram"
)

type Figure struct {
	NameMax    int
	DefaultX   int
	DefaultY   int
	DefaultPos nonogram.Pos
}
