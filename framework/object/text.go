package object

import (
	"github.com/simp7/nonograminGo/framework"
)

type Text interface {
	framework.Object
	CopyText() Text
}
