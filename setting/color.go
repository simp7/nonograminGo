package setting

import (
	"image/color"
)

type Color struct {
	Empty   color.RGBA
	Filled  color.RGBA
	Checked color.RGBA
	Wrong   color.RGBA
	Char    color.RGBA
}
