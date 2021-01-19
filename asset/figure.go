package asset

import "github.com/simp7/nonograminGo/util"

type Figure struct {
	NameMax    int
	WidthMax   int
	HeightMax  int
	DefaultX   int
	DefaultY   int
	DefaultPos util.Pos
}
