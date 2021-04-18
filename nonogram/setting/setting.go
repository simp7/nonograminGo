package setting

import (
	"github.com/simp7/nonograminGo/file"
	"github.com/simp7/nonograminGo/file/customPath"
	"github.com/simp7/nonograminGo/file/loader"
	"github.com/simp7/nonograminGo/file/updater"
	"github.com/simp7/nonograminGo/nonogram"
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
	updateLanguage(err)
}

func updateLanguage(err error) {
	if err != nil || !instance.Text.IsLatest("1.0") {
		updater.Language().Update()
		instance.Text, _ = text.New(instance.Language)
	}
}
