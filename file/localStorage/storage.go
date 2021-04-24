package localStorage

import (
	"github.com/simp7/nonograminGo/file"
)

type storage struct {
	path      customPath
	formatter file.Formatter
}

func newStorage(name PathName, formatter file.Formatter, leaf ...string) (*storage, error) {

	s := new(storage)
	var err error

	s.formatter = formatter
	path, err := Get(name)

	if err != nil {
		return nil, err
	}

	s.path = path.Append(leaf...)
	return s, err

}

func Language(language string, formatter file.Formatter) (*storage, error) {
	return newStorage(LANGUAGEDIR, formatter, language+".json")
}

func Map(name string, formatter file.Formatter) (*storage, error) {
	return newStorage(MAPSDIR, formatter, name+".nm")
}

func Setting(formatter file.Formatter) (*storage, error) {
	return newStorage(SETTING, formatter)
}

func (s storage) Save(i interface{}) error {
	save := &saver{s.path, s.formatter}
	return save.Save(i)
}

func (s storage) Load(i interface{}) error {
	load := &loader{s.path, s.formatter}
	return load.Load(i)
}
