package object

import (
	"github.com/simp7/nonograminGo/nonogram"
)

//Char is an basic unit of object that physically exist.
type Char interface {
	nonogram.Object
}
