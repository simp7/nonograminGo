package setting

import (
	"github.com/simp7/nonograminGo/errs"
	"github.com/simp7/nonograminGo/nonogram"
	"github.com/simp7/nonograminGo/nonogram/fileFormatter"
	"github.com/simp7/nonograminGo/nonogram/text"
	"os"
	"sync"
)

var instance *nonogram.Setting
var once sync.Once

func Get() *nonogram.Setting {

	once.Do(func() {

		instance = new(nonogram.Setting)
		pf := nonogram.GetPathFormatter()

		content, err := os.ReadFile(pf.GetPath("setting.json"))
		errs.Check(err)
		errs.Check(Load(content))

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

	pf := nonogram.GetPathFormatter()
	languageFile := pf.GetPath("language", instance.Language+".json")
	content, err := os.ReadFile(languageFile)
	errs.Check(err)

	instance.Text = text.New(content)
	if !instance.Text.IsLatest(version) {
		pf.UpdateLanguageFiles()
	}

}
