package loader

import (
	"github.com/simp7/nonograminGo/nonogram/file"
	"github.com/simp7/nonograminGo/nonogram/file/customPath"
	"github.com/simp7/nonograminGo/nonogram/file/formatter"
)

type loader struct {
	real      file.Path
	source    file.Path
	formatter file.Formatter
}

func Setting() *loader {
	m := new(loader)
	m.formatter = formatter.Json()
	m.real = customPath.SettingFile
	m.source = customPath.DefaultSettingFile
	return m
}

func Language(language string) *loader {
	m := new(loader)
	m.formatter = formatter.Json()
	m.real = customPath.LanguageFile(language)
	m.source = customPath.LanguageFile(language)
	return m
}

func (m *loader) load(from file.Path, target interface{}) {

	data, _ := file.ReadFile(from)

	m.formatter.GetRaw(data)
	m.formatter.Decode(&target)

}

func (m *loader) Load(target interface{}) {
	m.load(m.real, target)
}

func (m *loader) LoadDefault(target interface{}) {
	m.load(m.source, target)
}
