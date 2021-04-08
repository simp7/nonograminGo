package setting

import (
	"github.com/simp7/nonograminGo/nonogram/fileFormatter"
	"github.com/simp7/nonograminGo/util"
	"os"
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

		content, err := os.ReadFile(pf.GetPath("setting.json"))
		util.CheckErr(err)
		util.CheckErr(Load(content))

		updateLanguageFile("1.0")

	})

	return instance

}

func Load(content []byte) error {
	ff := fileFormatter.New()
	ff.GetRaw(content)
	return ff.Decode(&instance)
}

func updateLanguageFile(version string) {

	pf := util.GetPathFormatter()
	languageFile := pf.GetPath("language", instance.Language+".json")
	content, err := os.ReadFile(languageFile)
	util.CheckErr(err)

	instance.Text = NewText(content)
	if !instance.Text.IsLatest(version) {
		pf.UpdateLanguageFiles()
	}

}
