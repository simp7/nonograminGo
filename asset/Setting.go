package asset

import (
	"github.com/simp7/nonograminGo/util"
	"io/ioutil"
	"sync"
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
		ff := util.NewFileFormatter()

		instance.Color = defaultColor()
		instance.Figure = defaultFigure()

		instance.Language = "en.json"
		languageFile := pf.GetPath(instance.Language)
		content, _ := ioutil.ReadFile(languageFile)
		ff.GetRaw(string(content))

		util.CheckErr(ff.Decode(&instance.Text))

	})

	return instance

}
