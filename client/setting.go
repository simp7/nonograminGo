package client

import (
	"github.com/simp7/nonograminGo/file"
)

type Config struct {
	Color
	Text
	Figure
	Language string
}

func InitSetting(fs file.System, formatter file.Formatter) (*Config, error) {

	settingLoader, err := fs.Setting(formatter)
	if err != nil {
		return nil, err
	}

	config, err := loadSetting(settingLoader)
	if err != nil {
		return nil, err
	}

	languageLoader, err := fs.LanguageOf(config.Language, formatter)
	if err != nil {
		return nil, err
	}

	languageUpdater, err := fs.Language()
	if err != nil {
		return nil, err
	}

	text, err := loadText(languageLoader, languageUpdater)
	config.Text = text

	return config, err

}

func loadSetting(settingLoader file.Loader) (instance *Config, err error) {
	err = settingLoader.Load(&instance)
	return
}

func loadText(languageLoader file.Loader, languageUpdater file.Updater) (text Text, err error) {
	err = languageLoader.Load(&text)
	if err != nil || !text.IsLatest("1.0") {
		languageUpdater.Update()
		err = languageLoader.Load(&text)
	}
	return
}
