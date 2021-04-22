package localStorage

import (
	"github.com/simp7/nonograminGo/file"
	"github.com/simp7/nonograminGo/file/formatter"
)

type loader struct {
	path      file.Path
	formatter file.Formatter
}

func SettingLoader() (*loader, error) {

	m := new(loader)
	m.formatter = formatter.Json()

	var err error
	m.path, err = Get(SETTING)

	return m, err

}

func LanguageLoader(language string) (*loader, error) {

	m := new(loader)
	m.formatter = formatter.Json()

	languageDir, err := Get(LANGUAGEDIR)
	if err != nil {
		return nil, err
	}

	m.path = languageDir.Append(language + ".json")
	return m, err

}

func NonomapLoader(fileName string, formatter file.Formatter) (*loader, error) {

	m := new(loader)
	m.formatter = formatter

	mapsDir, err := Get(MAPSDIR)
	if err != nil {
		return nil, err
	}

	m.path = mapsDir.Append(fileName)
	return m, nil

}

func (m *loader) Load(target interface{}) error {

	data, err := ReadFile(m.path)
	if err != nil {
		return err
	}

	err = m.formatter.GetRaw(data)
	if err != nil {
		return err
	}

	return m.formatter.Decode(target)

}
