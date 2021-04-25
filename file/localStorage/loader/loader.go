package loader

import (
	"github.com/simp7/nonograminGo/file"
	"github.com/simp7/nonograminGo/file/localStorage"
	"github.com/simp7/nonograminGo/file/localStorage/customPath"
	"github.com/simp7/nonograminGo/file/localStorage/formatter"
)

type loader struct {
	path      file.Path
	formatter file.Formatter
}

func Setting() (*loader, error) {

	m := new(loader)
	m.formatter = formatter.Json()

	var err error
	m.path, err = customPath.Get(localStorage.SETTING)

	return m, err

}

func Language(language string) (*loader, error) {

	m := new(loader)
	m.formatter = formatter.Json()

	languageDir, err := customPath.Get(localStorage.LANGUAGEDIR)
	if err != nil {
		return nil, err
	}

	m.path = languageDir.Append(language + ".json")
	return m, err

}

func Nonomap(fileName string, formatter file.Formatter) (*loader, error) {

	m := new(loader)
	m.formatter = formatter

	mapsDir, err := customPath.Get(localStorage.MAPSDIR)
	if err != nil {
		return nil, err
	}

	m.path = mapsDir.Append(fileName)
	return m, nil

}

func (m *loader) Load(target interface{}) error {

	data, err := localStorage.ReadFile(m.path)

	if err != nil {
		return err
	}

	err = m.formatter.GetRaw(data)
	if err != nil {
		return err
	}

	return m.formatter.Decode(target)

}
