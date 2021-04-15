package saver

import (
	"github.com/simp7/nonograminGo/nonogram/file"
	"github.com/simp7/nonograminGo/nonogram/file/customPath"
)

type saver struct {
	path      file.Path
	formatter file.Formatter
}

func Nonomap(fileName string, formatter file.Formatter) *saver {
	s := new(saver)
	s.path = customPath.MapFile(fileName)
	s.formatter = formatter
	return s
}

func (s *saver) Save(i interface{}) error {
	s.formatter.Encode(i)
	return file.WriteFile(s.path, s.formatter.Content())
}
