package config

import (
	local2 "github.com/simp7/nonograminGo/file/localStorage"
	"github.com/simp7/nonograminGo/framework"
	"sync"
)

var instance *framework.Config
var once sync.Once

func Get() (*framework.Config, error) {

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

	rootUpdater, err := local2.AllUpdater()
	if err != nil {
		return err
	}

	rootUpdater.Update()
	return nil

}

func isFirst() bool {
	root, _ := local2.Get(local2.ROOT)
	return !local2.IsThere(root)
}

func load() error {

	err := loadSetting()
	if err != nil {
		return err
	}

	return loadText()

}

func loadSetting() error {

	settingLoader, err := local2.SettingLoader()
	if err != nil {
		return err
	}

	return settingLoader.Load(&instance)

}

func loadText() error {

	var err error
	instance.Text, err = New(instance.Language)

	return updateLanguage(err)

}

func updateLanguage(err error) error {

	if err != nil || !instance.Text.IsLatest("1.0") {

		languageUpdater, anotherErr := local2.LanguageUpdater()
		if anotherErr != nil {
			return anotherErr
		}

		languageUpdater.Update()
		instance.Text, anotherErr = New(instance.Language)
		return anotherErr

	}

	return nil

}
