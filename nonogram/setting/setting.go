package setting

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

func Get() *Setting {

	once.Do(func() {

		instance = new(Setting)
		pf := util.GetPathFormatter()

		content, err := ioutil.ReadFile(pf.GetPath("setting.json"))
		util.CheckErr(err)
		Load(content)

		languageFile := pf.GetPath("language", instance.Language+".json")
		content, err = ioutil.ReadFile(languageFile)
		util.CheckErr(err)
		instance.Text = NewText(content)

	})

	return instance

}

func Load(content []byte) {
	ff := util.NewFileFormatter()
	ff.GetRaw(content)
	ff.Decode(&instance)
}
