package localStorage

import (
	"fmt"
	"github.com/simp7/nonograminGo/file"
	"os"
	"strings"
)

type mapList struct {
	dirPath     []byte
	files       []os.DirEntry
	currentFile string
	order       int
}

func MapList() file.MapList {

	list := new(mapList)

	list.Refresh()
	list.order = 0

	return list

}

/*
	This function returns list of map whose number of maps are separated by 10.
	This function will be called when player enter the select page.
*/

func (l *mapList) Current() []string {

	list := make([]string, 10)

	for n := 0; n < 10; n++ {
		order := l.realIdx(n)
		if order < len(l.files) {
			fileName := l.files[order].Name()
			list[n] = fmt.Sprintf("%d. %s", n, trimSuffix(fileName))
		}
	}

	return list

}

/*
	This function gets player to the next page of list.
	This function will be called when player inputs left-arrow key.
*/

func (l *mapList) Next() {
	if 10*(l.order+1) >= len(l.files) {
		l.order = 0
	} else {
		l.order++
	}
}

/*
	This function gets player to the previous page of list
	This function will be called when player inputs right-arrow key.
*/

func (l *mapList) Prev() {
	if l.order == 0 {
		l.order = (len(l.files) - 1) / 10
	} else {
		l.order--
	}
}

/*
	This function returns player's current page.
	This function will be called with list of map, attached with list header.
*/

func (l *mapList) GetOrder() string {
	return fmt.Sprintf("(%d/%d)", l.order+1, len(l.files)/10+1)
}

/*
	This function gets nonomap data by number.
	This function will be called when user inputs number in select.
*/

func (l *mapList) GetMapName(idx int) (string, bool) {

	if idx >= len(l.files) {
		return "", false
	}

	l.currentFile = trimSuffix(l.files[l.realIdx(idx)].Name())
	return l.currentFile, true

}

func (l *mapList) GetCachedMapName() string {
	return l.currentFile
}

/*
	This function refresh list of map
	This function will be called after user create map so it contains added map.
*/

func (l *mapList) Refresh() error {

	mapDir, err := Get(MAPSDIR)
	if err != nil {
		return err
	}

	l.files, err = ReadDir(mapDir)
	return err

}

func (l *mapList) realIdx(idx int) int {
	return idx + 10*l.order
}

func trimSuffix(name string) string {
	return strings.TrimSuffix(name, ".nm")
}
