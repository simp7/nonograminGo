package setting

import (
	"github.com/simp7/nonograminGo/errs"
	"github.com/simp7/nonograminGo/nonogram"
	"github.com/simp7/nonograminGo/nonogram/file"
	"github.com/simp7/nonograminGo/nonogram/file/customPath"
	"github.com/simp7/nonograminGo/nonogram/file/loader"
	"github.com/simp7/nonograminGo/nonogram/file/updater"
	"github.com/simp7/nonograminGo/nonogram/text"
	"sync"
)

var instance *nonogram.Setting
var once sync.Once

func Get() *nonogram.Setting {

	once.Do(func() {
		if isFirst() {
			initializeDir()
		}
		load()
	})

	return instance

}

func initializeDir() {
	updater.All().Update()
}

func isFirst() bool {
	return !file.IsThere(customPath.Root)
}

func load() {
	loadSetting()
	loadText()
}

func loadSetting() {
	loader.Setting().Load(&instance)
}

func loadText() {
	var err error
	instance.Text, err = text.New(instance.Language)
	errs.Check(err)
}
