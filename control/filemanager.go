package control

import (
	"nonograminGo/asset"
	"nonograminGo/util"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

type FileManager struct {
	files       []os.FileInfo
	currentFile string
	order       int
}

func NewFileManager() *FileManager {
	fm := FileManager{}
	fm.currentFile = ""
	fm.order = 0

	var err error
	fm.files, err = ioutil.ReadDir("./maps")
	util.CheckErr(err)

	return &fm
}

/*
	This function returns list of map whose number of maps are seperated by 10.
	This function will be called when player enter the select page.
*/

func (fm *FileManager) GetMapList() []string {

	mapList := make([]string, 10)

	for n := 0; n < 10; n++ {
		if n+10*fm.order < len(fm.files) {
			mapList[n] = fmt.Sprintf("%d. %s", n, strings.TrimSuffix(fm.files[n+10*fm.order].Name(), ".nm"))
		}
	}

	return mapList

}

/*
	This function gets player to the next page of list.
	This function will be called when player inputs left-arrow key.
*/

func (fm *FileManager) NextList() {
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

func (fm *FileManager) PrevList() {
	if fm.order == 0 {
		fm.order = (len(fm.files) - 1) / 10
	} else {
		fm.order--
	}
}

/*
	This function returns player's current page.
	This function will be called with maplist, attached with list header.
*/

func (fm *FileManager) GetOrder() string {
	return fmt.Sprintf("(%d/%d)", fm.order+1, len(fm.files)/10+1)
}

/*
	This function gets nonomap data by number.
	This function will be called when user inputs number in select.
*/

func (fm *FileManager) GetMapDataByNumber(target int) string {

	if target >= len(fm.files) {
		return asset.StringMsgFileNotExist
	}
	fm.currentFile = fm.files[target+10*fm.order].Name()

	return fm.GetMapDataByName(fmt.Sprintf("./maps/%s", fm.currentFile))

}

/*
	This function gets nonomap data by name.
	This function will be called in GetMapDataByNumber.
*/

func (fm *FileManager) GetMapDataByName(target string) string {
	file, err := ioutil.ReadFile(target)
	util.CheckErr(err)

	return string(file)
}

/*
	This function returns file name without '.nm'.
	This function will be called with map list.
*/

func (fm *FileManager) GetCurrentMapName() string {
	return strings.TrimSuffix(fm.currentFile, ".nm")
}

/*
	This function creates file by bitmap that player generate in create mode.
	This function will be called when player finish create mode by pressing enter key.
*/

func (fm *FileManager) CreateMap(name string, width int, height int, bitmap [][]bool) {

	mapData := make([]int, height)
	nonoMapData := fmt.Sprintf("%d/%d", width, height)

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
		nonoMapData += fmt.Sprintf("/%d", v)
	}

	err := ioutil.WriteFile(fmt.Sprintf("./maps/%s.nm", name), []byte(nonoMapData), 0644)
	util.CheckErr(err)

}

/*
	This function refresh list of map
	This function will be called after user create map so it contains added map.
*/

func (fm *FileManager) RefreshMapList() {
	var err error
	fm.files, err = ioutil.ReadDir("./maps")
	util.CheckErr(err)
}
