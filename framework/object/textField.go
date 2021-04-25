package object

import (
	"github.com/simp7/nonograminGo/framework"
)

type TextField interface {
	framework.Object
	Activate()
	Deactivate()
}
