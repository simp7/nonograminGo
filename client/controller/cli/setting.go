package cli

import (
	"github.com/simp7/nonograminGo/client"
	"github.com/simp7/nonograminGo/file"
)

type Config struct {
	Color
	client.Text
	NameMax    int    //NameMax is an maximum length of name.
	DefaultPos Pos    //DefaultPos is an default position for display things.
	Language   string //Language is an language of this application.
}

func initSetting(fs file.System, formatter file.Formatter) (*Config, error) {

	settingLoader, err := fs.Setting(formatter)
	if err != nil {
		return nil, err
	}

	conf, err := loadSetting(settingLoader)
	if err != nil {
		return nil, err
	}

	languageLoader, err := fs.LanguageOf(conf.Language, formatter)
	if err != nil {
		return nil, err
	}

	languageUpdater, err := fs.Language()
	if err != nil {
		return nil, err
	}

	text, err := loadText(languageLoader, languageUpdater)
	conf.Text = text

	return conf, err

}

func loadSetting(settingLoader file.Loader) (instance *Config, err error) {
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
