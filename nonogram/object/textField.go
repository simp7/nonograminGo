package object

import (
	"github.com/simp7/nonograminGo/nonogram"
)

type TextField interface {
	nonogram.Object
	Activate()
	Deactivate()
}
