package localStorage

import (
	"github.com/simp7/nonograminGo/file"
	"sync"
)

var instance file.System
var once sync.Once

func Get() (file.System, error) {

	var err error

	once.Do(func() {
		if IsInitial() {
			var u file.Updater
			u, err = allUpdater()
			u.Update()
		}
		instance = new(system)
	})

	return instance, err

}

type system struct {
}

func (s *system) Map(name string, formatter file.Formatter) (file.Storage, error) {
	return mapStorage(name, formatter)
}

func (s *system) Setting(formatter file.Formatter) (file.Storage, error) {
	return settingStorage(formatter)
}

func (s *system) LanguageOf(language string, formatter file.Formatter) (file.Loader, error) {
	return languageLoader(language, formatter)
}

func (s *system) Language() (file.Updater, error) {
	return languageUpdater()
}

func (s *system) MapList() file.MapList {
	return newMapList()
}