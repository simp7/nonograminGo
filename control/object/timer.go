package object

import "github.com/simp7/nonograminGo/util"

type timer struct {
	object
	util.Timer
}

func NewTimer() Object {
	t := new(timer)
	return t
}
