package control

import (
	"../asset"
	"../util"
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

func (fm *FileManager) GetMapList() []string {

	mapList := make([]string, 10)

	for n := 0; n < 10; n++ {
		if n+10*fm.order < len(fm.files) {
			mapList[n] = fmt.Sprintf("%d. %s", n, strings.TrimSuffix(fm.files[n+10*fm.order].Name(), ".nm"))
		}
	}

	return mapList

}

func (fm *FileManager) NextList() {
	if 10*(fm.order+1) >= len(fm.files) {
		fm.order = 0
	} else {
		fm.order++
	}
}

func (fm *FileManager) PrevList() {
	if fm.order == 0 {
		fm.order = (len(fm.files) - 1) / 10
	} else {
		fm.order--
	}
}

func (fm *FileManager) GetCurrentOrder() int {
	return fm.order + 1
}

func (fm *FileManager) GetMaxOrder() int {
	return len(fm.files)/10 + 1
}

func (fm *FileManager) GetMapDataByNumber(target int) string {

	if target >= len(fm.files) {
		return asset.StringMsgFileNotExist
	}
	fm.currentFile = fm.files[target].Name()

	return fm.GetMapDataByName(fmt.Sprintf("./maps/%s", fm.currentFile))

}

func (fm *FileManager) GetMapDataByName(target string) string {
	file, err := ioutil.ReadFile(target)
	util.CheckErr(err)

	return string(file)
}

func (fm *FileManager) GetCurrentMapName() string {
	return strings.TrimSuffix(fm.currentFile, ".nm")
}

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

func (fm *FileManager) RefreshMapList() {
	var err error
	fm.files, err = ioutil.ReadDir("./maps")
	util.CheckErr(err)
}
