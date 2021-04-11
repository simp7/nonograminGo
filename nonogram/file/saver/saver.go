package saver

import (
	"github.com/simp7/nonograminGo/nonogram"
	"github.com/simp7/nonograminGo/nonogram/file"
	"github.com/simp7/nonograminGo/nonogram/file/customPath"
	"github.com/simp7/nonograminGo/nonogram/file/formatter"
)

type saver struct {
	path      file.Path
	formatter file.Formatter
}

func Nonomap(fileName string, prototype nonogram.Map) *saver {
	s := new(saver)
	s.formatter = formatter.Map(prototype)
	s.path = customPath.MapFile(fileName)
	return s
}

func (s *saver) Save(i interface{}) {
	s.formatter.Encode(i)
	file.WriteFile(s.path, s.formatter.Content())
}
