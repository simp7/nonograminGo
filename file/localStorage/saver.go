package localStorage

import (
	"github.com/simp7/nonograminGo/file"
)

type saver struct {
	path      customPath
	formatter file.Formatter
}

func (s *saver) Save(i interface{}) error {

	err := s.formatter.Encode(i)
	if err != nil {
		return err
	}

	return WriteFile(s.path, s.formatter.Content())

}
