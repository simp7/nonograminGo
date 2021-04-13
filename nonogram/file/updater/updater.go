package updater

import (
	"embed"
	_ "embed"
	"github.com/simp7/nonograminGo/errs"
	"github.com/simp7/nonograminGo/nonogram/file"
	"github.com/simp7/nonograminGo/nonogram/file/customPath"
	"path"
)

//go:embed skel
var f embed.FS

type updater struct {
	source string
	target file.Path
}

func new(source string, target file.Path) *updater {
	return &updater{source: source, target: target}
}

func All() *updater {
	return new("skel", customPath.Root)
}

func Language() *updater {
	return new(path.Join("skel", "language"), customPath.LanguageDir)
}

func (u *updater) Update() {
	u.updateDir(u.source, u.target)
}

func (u *updater) updateDir(from string, to file.Path) {

	file.MkDir(to)
	files, _ := f.ReadDir(from)

	for _, v := range files {

		source := path.Join(from, v.Name())
		target := to.Append(v.Name())

		if v.IsDir() {
			u.updateDir(source, target)
		} else {
			data, _ := f.ReadFile(source)
			errs.Check(file.WriteFile(target, data))
		}

	}

}
