package updater

import (
	"embed"
	"github.com/simp7/nonograminGo/file"
	"github.com/simp7/nonograminGo/file/localStorage"
	"github.com/simp7/nonograminGo/file/localStorage/customPath"
	"path"
)

//go:embed skel
var f embed.FS

type updater struct {
	source string
	target file.Path
}

func new(source string, target localStorage.PathName) (*updater, error) {
	path, err := customPath.Get(target)
	return &updater{source: source, target: path}, err
}

func All() (*updater, error) {
	return new("skel", localStorage.ROOT)
}

func Language() (*updater, error) {
	return new(path.Join("skel", "language"), localStorage.LANGUAGEDIR)
}

func (u *updater) Update() {
	u.updateDir(u.source, u.target)
}

func (u *updater) updateDir(from string, to file.Path) error {

	err := localStorage.MkDir(to)
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
			err = localStorage.WriteFile(target, data)
		}

	}

	return err

}
