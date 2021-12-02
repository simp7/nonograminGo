package main

import (
	"fmt"
	"strings"
)

type mapList struct {
	unit        int
	files       []string
	currentFile string
	currentPage int
}

func newMapList(list []string) *mapList {
	m := new(mapList)
	m.unit = 10
	m.files = list
	m.currentPage = 1
	return m
}

func (l *mapList) Current() []string {

	list := make([]string, l.unit)

	for n := 0; n < l.unit; n++ {
		order := l.realIdx(n)
		if order < len(l.files) {
			fileName := l.files[order]
			list[n] = fmt.Sprintf("%d. %s", n, trimSuffix(fileName))
		}
	}

	return list

}

func (l *mapList) Next() {
	if l.unit*l.currentPage >= len(l.files) {
		l.currentPage = 1
		return
	}
	l.currentPage++
}

func (l *mapList) Prev() {
	if l.currentPage == 1 {
		l.currentPage = l.LastPage()
		return
	}
	l.currentPage--
}

func (l *mapList) CurrentPage() int {
	return l.currentPage
}

func (l *mapList) LastPage() int {
	return (len(l.files)-1)/l.unit + 1
}

func (l *mapList) GetMapName(idx int) (string, bool) {

	if l.realIdx(idx) >= len(l.files) {
		return "", false
	}

	l.currentFile = trimSuffix(l.files[l.realIdx(idx)])
	return l.currentFile, true

}

func (l *mapList) GetCachedMapName() string {
	return l.currentFile
}

func (l *mapList) realIdx(idx int) int {
	return idx + l.unit*(l.currentPage-1)
}

func trimSuffix(name string) string {
	return strings.TrimSuffix(name, ".nm")
}

func (l *mapList) IsEmpty() bool {
	return len(l.files) == 0
}
