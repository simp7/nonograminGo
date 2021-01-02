package util

import (
	"path"
	"sync"
)

type PathFormatter interface {
	GetPath(of ...string) string
}

type pathFormatter struct {
}

var instance PathFormatter
var once sync.Once

func GetPathFormatter() PathFormatter {

	once.Do(func() {
		instance = newPathFormatter()
	})

	return instance

}

func newPathFormatter() PathFormatter {

	p := new(pathFormatter)
	return p

}

func (p *pathFormatter) GetPath(target ...string) string {

	current := "."

	for _, element := range target {
		current = path.Join(current, element)
	}

	return current

}
