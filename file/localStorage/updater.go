package localStorage

import (
	"embed"
	"github.com/simp7/nonograminGo/file"
	"path"
)

//go:embed skel
var f embed.FS

type updater struct {
	source string
	target file.Path
}

func newUpdater(source string, target PathName) (*updater, error) {
	path, err := Get(target)
	return &updater{source: source, target: path}, err
}

func AllUpdater() (*updater, error) {
	return newUpdater("skel", ROOT)
}

func LanguageUpdater() (*updater, error) {
	return newUpdater(path.Join("skel", "language"), LANGUAGEDIR)
}

func (u *updater) Update() {
	u.updateDir(u.source, u.target)
}

func (u *updater) updateDir(from string, to file.Path) error {

	err := MkDir(to)
	if err != nil {
		return err
	}

	files, _ := f.ReadDir(from)

	for _, v := range files {

		source := path.Join(from, v.Name())
		target := to.Append(v.Name())

		if v.IsDir() {
			err = u.updateDir(source, target)
		} else {
			data, _ := f.ReadFile(source)
			err = WriteFile(target, data)
		}

	}

	return err

}
