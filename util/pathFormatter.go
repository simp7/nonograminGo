package util

import (
	"embed"
	"errors"
	"os"
	"path"
	"sync"
)

//go:embed default_setting.json
//go:embed language
//go:embed default_maps
var f embed.FS

type PathFormatter interface {
	GetPath(of ...string) string
	UpdateLanguageFiles()
}

type pathFormatter struct {
	root string
}

var instance PathFormatter
var once sync.Once

func GetPathFormatter() PathFormatter {

	once.Do(func() {

		workingDir, ok := os.LookupEnv("HOME")
		if !ok {
			CheckErr(errors.New("HOME not exist"))
		}

		workingDir = path.Join(workingDir, "nonogram")
		instance = newPathFormatter(workingDir)

		_, err := os.ReadDir(workingDir)
		if err != nil {
			os.Mkdir(workingDir, 0755)
			initialize()
		}

	})

	return instance

}

func newPathFormatter(root string) PathFormatter {

	p := new(pathFormatter)
	p.root = root

	return p

}

func initialize() {
	initDefaultMap()
	initDefaultSetting()
	initLanguage()
}

func initDefaultSetting() {
	copyFile("default_setting.json", instance.GetPath("setting.json"))
}

func initDefaultMap() {
	os.Mkdir(instance.GetPath("maps"), 0755)
	copyDir("default_maps", "maps")
}

func initLanguage() {
	os.Mkdir(instance.GetPath("language"), 0755)
	copyDir("language", "language")
}

func copyDir(from string, to string) {
	files, _ := f.ReadDir(from)
	for _, file := range files {
		copyFile(from+string(os.PathSeparator)+file.Name(), instance.GetPath(to, file.Name()))
	}
}

func copyFile(from string, to string) {
	data, _ := f.ReadFile(from)
	_ = os.Remove(to)
	CheckErr(os.WriteFile(to, data, 0644))
}

func (p *pathFormatter) UpdateLanguageFiles() {
	initLanguage()
}

func (p *pathFormatter) GetPath(target ...string) string {

	current := p.root

	for _, element := range target {
		current = path.Join(current, element)
	}

	return current

}

func (p *pathFormatter) moveFile(fileName, from, to string) {

	f := path.Join(from, fileName)
	t := path.Join(to, fileName)

	CheckErr(os.Rename(f, t))

}

func (p *pathFormatter) MoveBase(to string) {

	list := getAllFilesFrom(p.root)

	for _, name := range list {
		p.moveFile(name, p.root, to)
	}

	p.root = to

}

//getAllFilesFrom gets and returns absolute path.
func getAllFilesFrom(parentDirectory string) []string {

	result := make([]string, 0)
	files, _ := os.ReadDir(parentDirectory)

	for _, file := range files {
		if file.IsDir() {
			inner := getAllFilesFrom(parentDirectory + file.Name())
			for i := range inner {
				inner[i] = path.Join(inner[i], file.Name())
			}
			result = append(result, inner...)
		} else {
			result = append(result, parentDirectory+file.Name())
		}
	}

	return result

}
