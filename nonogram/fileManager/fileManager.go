package fileManager

import (
	"fmt"
	"github.com/simp7/nonograminGo/nonogram"
	"github.com/simp7/nonograminGo/nonogram/fileFormatter"
	"github.com/simp7/nonograminGo/nonogram/nonomap"
	"github.com/simp7/nonograminGo/nonogram/setting"
	"github.com/simp7/nonograminGo/util"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

type fileManager struct {
	dirPath     []byte
	files       []os.FileInfo
	currentFile string
	order       int
	util.PathFormatter
	nonogram.FileFormatter
	*setting.Setting
}

func New() nonogram.FileManager {

	fm := new(fileManager)
	fm.order = 0
	fm.PathFormatter = util.GetPathFormatter()
	fm.FileFormatter = fileFormatter.Map()

	var err error

	fm.files, err = ioutil.ReadDir(fm.GetPath("maps"))
	util.CheckErr(err)
	fm.Setting = setting.Get()

	return fm

}

/*
	This function returns list of map whose number of maps are separated by 10.
	This function will be called when player enter the select page.
*/

func (fm *fileManager) GetMapList() []string {

	mapList := make([]string, 10)

	for n := 0; n < 10; n++ {
		if n+10*fm.order < len(fm.files) {
			mapList[n] = fmt.Sprintf("%d. %s", n, strings.TrimSuffix(fm.files[n+10*fm.order].Name(), ".nm"))
		}
	}

	return mapList

} //TODO: Separate suffix formatting function.

/*
	This function gets player to the next page of list.
	This function will be called when player inputs left-arrow key.
*/

func (fm *fileManager) NextList() {
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

func (fm *fileManager) PrevList() {
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

func (fm *fileManager) GetOrder() string {
	return fmt.Sprintf("(%d/%d)", fm.order+1, len(fm.files)/10+1)
}

/*
	This function gets nonomap data by number.
	This function will be called when user inputs number in select.
*/

func (fm *fileManager) GetMapDataByNumber(target int) (nonogram.Map, bool) {

	if target >= len(fm.files) {
		return nil, false
	}
	fm.currentFile = fm.files[target+10*fm.order].Name()

	return fm.GetMapDataByName(fm.GetPath("maps", fm.currentFile))

}

/*
	This function gets nonomap data by name.
	This function will be called in GetMapDataByNumber.
*/

func (fm *fileManager) GetMapDataByName(target string) (nonogram.Map, bool) {

	file, err := ioutil.ReadFile(target)
	util.CheckErr(err)

	fm.GetRaw(file)
	result := nonomap.New()
	err = fm.Decode(result)
	util.CheckErr(err)

	return result, true

}

/*
	This function returns file name without '.nm'.
	This function will be called with map list.
*/

func (fm *fileManager) GetCurrentMapName() string {
	return strings.TrimSuffix(fm.currentFile, ".nm")
}

/*
	This function creates file by bitmap that player generate in create mode.
	This function will be called when player finish create mode by pressing enter key.
*/

func (fm *fileManager) CreateMap(name string, width int, height int, bitmap [][]bool) {

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

	err := ioutil.WriteFile(fm.GetPath("maps", name+".nm"), []byte(nonoMapData), 0644)
	util.CheckErr(err)

}

/*
	This function refresh list of map
	This function will be called after user create map so it contains added map.
*/

func (fm *fileManager) RefreshMapList() {
	var err error
	fm.files, err = ioutil.ReadDir(fm.GetPath("maps"))
	util.CheckErr(err)
}
