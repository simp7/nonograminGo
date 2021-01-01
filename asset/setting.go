package asset

import (
	"github.com/simp7/nonograminGo/util"
	"io/ioutil"
	"sync"
)

const (
	EN string = "en"
	KR string = "kr"
)

type Setting struct {
	Color
	Text
	Figure
	Language string
}

var instance *Setting
var once sync.Once

func GetSetting() *Setting {

	once.Do(func() {

		instance = new(Setting)
		pf := util.GetPathFormatter()

		instance.Color = defaultColor()
		instance.Figure = defaultFigure()

		instance.Language = EN
		languageFile := pf.GetPath("asset", "language", instance.Language+".json")
		content, _ := ioutil.ReadFile(languageFile)
		instance.Text = NewText(content)

	})

	return instance

}
