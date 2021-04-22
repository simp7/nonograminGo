package localStorage

import (
	"github.com/simp7/nonograminGo/file"
)

type saver struct {
	path      file.Path
	formatter file.Formatter
}

func NonomapSaver(fileName string, formatter file.Formatter) (*saver, error) {

	s := new(saver)

	mapDir, err := Get(MAPSDIR)
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

	return WriteFile(s.path, s.formatter.Content())

}
