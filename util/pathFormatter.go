package util

import (
	"os"
	"path"
	"sync"
)

type PathFormatter interface {
	GetPath(of ...string) string
}

type pathFormatter struct {
	base string
}

var instance PathFormatter
var once sync.Once

func GetPathFormatter() PathFormatter {

	once.Do(func() {
		workingDir, err := os.Getwd()
		CheckErr(err)
		instance = newPathFormatter(workingDir)
	})

	return instance

}

func newPathFormatter(base string) PathFormatter {

	p := new(pathFormatter)
	p.base = base

	return p

}

func (p *pathFormatter) GetPath(target ...string) string {

	current := p.base

	for _, element := range target {
		current = path.Join(current, element)
	}

	return current

}

func (p *pathFormatter) MoveBase(to string) {

	list := getAllFileNames(p.base)

	for _, name := range list {
		CheckErr(os.Rename(path.Join(p.base, name), path.Join(to, name)))
	}

	p.base = to

}

func getAllFileNames(filePath string) []string {

	result := make([]string, 0)
	files, _ := os.ReadDir(filePath)

	for _, file := range files {
		if file.IsDir() {
			inner := getAllFileNames(file.Name())
			for i := range inner {
				inner[i] = path.Join(inner[i], file.Name())
			}
			result = append(result, inner...)
		} else {
			result = append(result, file.Name())
		}
	}

	return result

}
