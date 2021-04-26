package config

import (
	"github.com/simp7/nonograminGo/client"
	"github.com/simp7/nonograminGo/file/formatter"
	"github.com/simp7/nonograminGo/file/localStorage"
	"sync"
)

var instance *client.Config
var once sync.Once

func Get() (*client.Config, error) {

	var err error

	once.Do(func() {
		if localStorage.IsInitial() {
			err = initializeDir()
			if err != nil {
				return
			}
		}
		err = load()
	})

	return instance, err

}

func initializeDir() error {

	rootUpdater, err := localStorage.AllUpdater()
	if err != nil {
		return err
	}

	rootUpdater.Update()
	return nil

}

func load() error {

	err := loadSetting()
	if err != nil {
		return err
	}

	return loadText()

}

func loadSetting() error {

	settingLoader, err := localStorage.Setting(formatter.Json())
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

		languageUpdater, anotherErr := localStorage.LanguageUpdater()
		if anotherErr != nil {
			return anotherErr
		}

		languageUpdater.Update()
		instance.Text, anotherErr = New(instance.Language)
		return anotherErr

	}

	return nil

}
