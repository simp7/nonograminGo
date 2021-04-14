package loader

import (
	"fmt"
	"github.com/simp7/nonograminGo/errs"
	"github.com/simp7/nonograminGo/nonogram/file"
	"github.com/simp7/nonograminGo/nonogram/file/customPath"
	"math"
	"os"
	"strings"
)

type mapList struct {
	dirPath     []byte
	files       []os.DirEntry
	currentFile string
	order       int
}

func New() file.MapList {

	list := new(mapList)

	list.refresh()
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
		if n+10*l.order < len(l.files) {
			list[n] = fmt.Sprintf("%d. %s", n, strings.TrimSuffix(l.files[n+10*l.order].Name(), ".nm"))
		}
	}

	return list

} //TODO: Separate suffix formatting function.

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

func (l *mapList) GetMapName(target int) (string, bool) {

	if target >= len(l.files) {
		return "", false
	}

	l.currentFile = l.files[target+10*l.order].Name()
	return l.currentFile, true

}

func (l *mapList) GetCachedMapName() string {
	return l.currentFile
}

/*
	This function creates file by bitmap that player generate in create mode.
	This function will be called when player finish create mode by pressing enter key.
*/

func (l *mapList) CreateMap(name string, width int, height int, bitmap [][]bool) {

	mapData := make([]int, height)
	nonomapData := fmt.Sprintf("%d/%d", width, height)

	for n := range bitmap {
		result := 0
		for m, v := range bitmap[n] {
			if v {
				result += int(math.Pow(2, float64(width-m-1)))
			}
		}
		mapData[n] = result
	}

	for _, v := range mapData {
		nonomapData += fmt.Sprintf("/%d", v)
	}

	err := file.WriteFile(customPath.MapFile(name), []byte(nonomapData))
	errs.Check(err)

	l.refresh()

}

/*
	This function refresh list of map
	This function will be called after user create map so it contains added map.
*/

func (l *mapList) refresh() {
	var err error
	l.files, err = file.ReadDir(customPath.MapsDir)
	errs.Check(err)
}
