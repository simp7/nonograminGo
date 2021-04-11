package loader

import (
	"embed"
	"fmt"
	"github.com/simp7/nonograminGo/errs"
	"github.com/simp7/nonograminGo/nonogram"
	"github.com/simp7/nonograminGo/nonogram/file"
	"github.com/simp7/nonograminGo/nonogram/file/customPath"
	"math"
	"os"
	"strings"
)

//go embed:skel
var f embed.FS

type mapList struct {
	dirPath      []byte
	files        []os.DirEntry
	currentFile  string
	order        int
	mapPrototype nonogram.Map
}

func New(mapPrototype nonogram.Map) file.MapList {

	fm := new(mapList)
	fm.order = 0

	var err error

	fm.files, err = file.ReadDir(customPath.MapsDir)
	fm.mapPrototype = mapPrototype
	errs.Check(err)

	return fm

}

/*
	This function returns list of map whose number of maps are separated by 10.
	This function will be called when player enter the select page.
*/

func (fm *mapList) GetAll() []string {

	list := make([]string, 10)

	for n := 0; n < 10; n++ {
		if n+10*fm.order < len(fm.files) {
			list[n] = fmt.Sprintf("%d. %s", n, strings.TrimSuffix(fm.files[n+10*fm.order].Name(), ".nm"))
		}
	}

	return list

} //TODO: Separate suffix formatting function.

/*
	This function gets player to the next page of list.
	This function will be called when player inputs left-arrow key.
*/

func (fm *mapList) Next() {
	if 10*(fm.order+1) >= len(fm.files) {
		fm.order = 0
	} else {
		fm.order++
	}
}

/*
	This function gets player to the previous page of list
	This function will be called when player inputs right-arrow key.
*/

func (fm *mapList) Prev() {
	if fm.order == 0 {
		fm.order = (len(fm.files) - 1) / 10
	} else {
		fm.order--
	}
}

/*
	This function returns player's current page.
	This function will be called with list of map, attached with list header.
*/

func (fm *mapList) GetOrder() string {
	return fmt.Sprintf("(%d/%d)", fm.order+1, len(fm.files)/10+1)
}

/*
	This function gets nonomap data by number.
	This function will be called when user inputs number in select.
*/

func (fm *mapList) GetMapName(target int) (string, bool) {

	if target >= len(fm.files) {
		return "", false
	}

	fm.currentFile = fm.files[target+10*fm.order].Name()
	return fm.currentFile, true

}

func (fm *mapList) GetCachedMapName() string {
	return fm.currentFile
}

/*
	This function creates file by bitmap that player generate in create mode.
	This function will be called when player finish create mode by pressing enter key.
*/

func (fm *mapList) CreateMap(name string, width int, height int, bitmap [][]bool) {

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

}

/*
	This function refresh list of map
	This function will be called after user create map so it contains added map.
*/

func (fm *mapList) Refresh() {
	var err error
	fm.files, err = file.ReadDir(customPath.MapsDir)
	errs.Check(err)
}

//func initialize() {
//	initDefaultMap()
//	initDefaultSetting()
//	initLanguage()
//}
//
//func initDefaultSetting() {
//	install(customPath.DefaultSettingFile, customPath.SettingFile)
//}
//
//func initDefaultMap() {
//	install(customPath.DefaultMapsDir, customPath.MapsDir)
//}
//
//func initLanguage() {
//	install(customPath.DefaultLanguageDir, customPath.LanguageDir)
//}
//
//func statOf(target file.Path) fs.FileInfo {
//
//	opened, err := f.Open(target.String())
//	errs.Check(err)
//
//	stat, err := opened.Stat()
//	errs.Check(err)
//
//	return stat
//
//}
//
//func install(from file.Path, to file.Path) {
//
//	source := from.String()
//	target := to.String()
//	stat := statOf(from)
//
//	if stat.IsDir() {
//		files, _ := f.ReadDir(source)
//		for _, v := range files {
//			install(from.Append(v.Name()), to.Append(v.Name()))
//		}
//		file.MkDir(to)
//	} else {
//		_ = os.Remove(target)
//		data, _ := f.ReadFile(source)
//		errs.Check(file.WriteFile(to, data))
//	}
//
//}
