package saver

import (
	"github.com/simp7/nonograminGo/file"
	"github.com/simp7/nonograminGo/file/localStorage"
	customPath "github.com/simp7/nonograminGo/file/localStorage/customPath"
)

type saver struct {
	path      file.Path
	formatter file.Formatter
}

func Nonomap(fileName string, formatter file.Formatter) (*saver, error) {

	s := new(saver)

	mapDir, err := customPath.Get(localStorage.MAPSDIR)
	if err != nil {
		return nil, err
	}

	s.path = mapDir.Append(fileName + ".nm")
	s.formatter = formatter

	return s, nil

}

func (s *saver) Save(i interface{}) error {

	err := s.formatter.Encode(i)
	if err != nil {
		return err
	}

	return localStorage.WriteFile(s.path, s.formatter.Content())

}
