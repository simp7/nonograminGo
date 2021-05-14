package cli

import (
	"github.com/simp7/nonograminGo/client"
	"github.com/simp7/nonograminGo/file"
)

type config struct {
	Color
	client.Text
	Figure
	Language string
}

func InitSetting(fs file.System, formatter file.Formatter) (*config, error) {

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

func loadSetting(settingLoader file.Loader) (instance *config, err error) {
	err = settingLoader.Load(&instance)
	return
}

func loadText(languageLoader file.Loader, languageUpdater file.Updater) (text client.Text, err error) {
	text = new(TextData)
	err = languageLoader.Load(&text)
	if err != nil || !text.IsLatest("1.0") {
		languageUpdater.Update()
		err = languageLoader.Load(&text)
	}
	return
}
