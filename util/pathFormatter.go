package util

import (
	"errors"
	"os"
	"path"
	"sync"
)

type PathFormatter interface {
	GetPath(of ...string) string
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
		os.Mkdir(workingDir, 755)
		instance = newPathFormatter(workingDir)
	})

	return instance

}

func newPathFormatter(root string) PathFormatter {

	p := new(pathFormatter)
	p.root = root

	return p

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
			result = append(result, parentDirectory, file.Name())
		}
	}

	return result

}
