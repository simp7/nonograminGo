package localStorage

import (
	"embed"
	"path"
)

//go:embed skel
var f embed.FS

type updater struct {
	source string
	target customPath
}

func newUpdater(source string, target PathName) (*updater, error) {
	path, err := get(target)
	return &updater{source: source, target: path}, err
}

func allUpdater() (*updater, error) {
	return newUpdater("skel", ROOT)
}

func languageUpdater() (*updater, error) {
	return newUpdater(path.Join("skel", "language"), LANGUAGEDIR)
}

func (u *updater) Update() {
	u.updateDir(u.source, u.target)
}

func (u *updater) updateDir(from string, to customPath) error {

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
