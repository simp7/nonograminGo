package setting

import (
	"github.com/simp7/nonograminGo/file/localStorage"
	"github.com/simp7/nonograminGo/file/localStorage/customPath"
	"github.com/simp7/nonograminGo/file/localStorage/loader"
	"github.com/simp7/nonograminGo/file/localStorage/updater"
	"github.com/simp7/nonograminGo/framework"
	"github.com/simp7/nonograminGo/framework/text"
	"sync"
)

var instance *framework.Setting
var once sync.Once

func Get() (*framework.Setting, error) {

	var err error

	once.Do(func() {
		if isFirst() {
			err = initializeDir()
			if err != nil {
				return
			}
		}
		load()
	})

	return instance, nil

}

func initializeDir() error {

	rootUpdater, err := updater.All()
	if err != nil {
		return err
	}

	rootUpdater.Update()
	return nil

}

func isFirst() bool {
	root, _ := customPath.Get(localStorage.ROOT)
	return !localStorage.IsThere(root)
}

func load() error {

	err := loadSetting()
	if err != nil {
		return err
	}

	return loadText()

}

func loadSetting() error {

	settingLoader, err := loader.Setting()
	if err != nil {
		return err
	}

	return settingLoader.Load(&instance)

}

func loadText() error {

	var err error
	instance.Text, err = text.New(instance.Language)

	return updateLanguage(err)

}

func updateLanguage(err error) error {

	if err != nil || !instance.Text.IsLatest("1.0") {

		languageUpdater, anotherErr := updater.Language()
		if anotherErr != nil {
			return anotherErr
		}

		languageUpdater.Update()
		instance.Text, anotherErr = text.New(instance.Language)
		return anotherErr

	}

	return nil

}
