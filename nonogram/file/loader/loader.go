package loader

import (
	"github.com/simp7/nonograminGo/errs"
	"github.com/simp7/nonograminGo/nonogram/file"
	"github.com/simp7/nonograminGo/nonogram/file/customPath"
	"github.com/simp7/nonograminGo/nonogram/file/formatter"
)

type loader struct {
	path      file.Path
	formatter file.Formatter
}

func Setting() *loader {
	m := new(loader)
	m.formatter = formatter.Json()
	m.path = customPath.SettingFile
	return m
}

func Language(language string) *loader {
	m := new(loader)
	m.formatter = formatter.Json()
	m.path = customPath.LanguageFile(language)
	return m
}

func Nonomap(fileName string, formatter file.Formatter) *loader {
	m := new(loader)
	m.formatter = formatter
	m.path = customPath.MapFile(fileName)
	return m
}

func (m *loader) Load(target interface{}) error {

	data, err := file.ReadFile(m.path)

	m.formatter.GetRaw(data)
	errs.Check(m.formatter.Decode(target))

	return err

}
