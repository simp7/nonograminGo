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
	base string
}

var instance PathFormatter
var once sync.Once

func GetPathFormatter() PathFormatter {

	once.Do(func() {
		workingDir, ok := os.LookupEnv("GOPATH")
		if !ok {
			CheckErr(errors.New("GOPATH not exist"))
		}
		workingDir = path.Join(workingDir, "src", "github.com", "simp7", "nonograminGo")
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

func (p *pathFormatter) moveFile(fileName, from, to string) {

	f := path.Join(from, fileName)
	t := path.Join(to, fileName)

	CheckErr(os.Rename(f, t))

}

func (p *pathFormatter) MoveBase(to string) {

	list := getAllFileNames(p.base)

	for _, name := range list {
		p.moveFile(name, p.base, to)
	}

	p.base = to

}

//TODO: when selected path contains directory, shown inner file should include that directory.
func getAllFileNames(filePath string) []string {

	result := make([]string, 0)
	files, _ := os.ReadDir(filePath)

	for _, file := range files {
		if file.IsDir() {
			inner := getAllFileNames(filePath + file.Name())
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
