package object

import (
	"github.com/simp7/nonograminGo/nonogram"
)

type Text interface {
	nonogram.Object
	CopyText() Text
}
