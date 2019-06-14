package control

import (
	"../util"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type FileManager struct {
	files       []os.FileInfo
	currentFile string
}

func NewFileManager() *FileManager {
	fm := FileManager{}
	fm.currentFile = ""

	var err error
	fm.files, err = ioutil.ReadDir("./maps")
	util.CheckErr(err)

	return &fm
}

func (fm *FileManager) GetMapList() []string {
	mapList := []string{}

	for n, file := range fm.files {
		mapList = append(mapList, fmt.Sprintf("%d. %s", n+1, strings.TrimSuffix(file.Name(), ".nm")))
	}

	return mapList
}

func (fm *FileManager) GetMapDataByNumber(target int) string {
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
