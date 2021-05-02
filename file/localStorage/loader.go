package localStorage

import (
	"github.com/simp7/nonograminGo/file"
)

type loader struct {
	path      customPath
	formatter file.Formatter
}

func (l *loader) Load(target interface{}) error {

	data, err := readFile(l.path)
	if err != nil {
		return err
	}

	err = l.formatter.GetRaw(data)
	if err != nil {
		return err
	}

	return l.formatter.Decode(target)

}

func languageLoader(language string, formatter file.Formatter) (*loader, error) {

	path, err := get(LANGUAGEDIR)

	if err != nil {
		return nil, err
	}

	return &loader{path.Append(language + ".json"), formatter}, nil

}
