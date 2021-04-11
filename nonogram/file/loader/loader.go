package loader

import (
	"github.com/simp7/nonograminGo/nonogram"
	"github.com/simp7/nonograminGo/nonogram/file"
	"github.com/simp7/nonograminGo/nonogram/file/customPath"
	"github.com/simp7/nonograminGo/nonogram/file/formatter"
)

type loader struct {
	path      file.Path
	source    file.Path
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

func Nonomap(fileName string, prototype nonogram.Map) *loader {
	m := new(loader)
	m.formatter = formatter.Map(prototype)
	m.path = customPath.MapFile(fileName)
	return m
}

func (m *loader) Load(target interface{}) error {

	data, err := file.ReadFile(m.path)

	m.formatter.GetRaw(data)
	m.formatter.Decode(&target)

	return err

}
