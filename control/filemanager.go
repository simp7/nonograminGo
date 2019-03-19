package control

import (
	"../util"
	"fmt"
	"io/ioutil"
	"os"
)

type FileManager struct {
	files []os.FileInfo
}

func NewFileManager() *FileManager {
	fm := FileManager{}
	var err error
	fm.files, err = ioutil.ReadDir("./maps")
	util.CheckErr(err)
	return &fm
}

func (fm *FileManager) GetMapList() []string {
	mapList := []string{}

	for n, file := range fm.files {
		mapList = append(mapList, fmt.Sprintf("%d. %s", n+1, file.Name()))
	}

	return mapList
}

func (fm *FileManager) GetMapDataByNumber(target int) string {
	return fm.GetMapDataByName(fm.files[target].Name())
}

func (fm *FileManager) GetMapDataByName(target string) string {
	file, err := ioutil.ReadFile(target)
	util.CheckErr(err)
	return string(file)
}
