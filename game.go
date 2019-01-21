package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type nonomap struct {
	width   int
	height  int
	mapdata []int //each elements of mapdata contains data of 32 blocks. It is designated by bitmap, which 1 is filled and 0 is blank
}

func newNonomap(fileName string) *nonomap {

	var imported nonomap

	file, err := ioutil.ReadFile(fmt.Sprintf("%s.nm", fileName))
	CheckErr(err)
	elements := strings.Split(string(file), "/")
	//Extract all data from wanted file.

	imported.width, err = strconv.Atoi(elements[0])
	imported.height, err = strconv.Atoi(elements[1])
	CheckErr(err)
	//Extract map's size from file

	for _, v := range elements[2:] {
		temp, err := strconv.Atoi(v)
		imported.mapdata = append(imported.mapdata, temp)
		CheckErr(err)
	}
	//Extract map's answer from file.

	for _, v := range imported.mapdata {
		if float64(v) >= math.Pow(2, float64(imported.width)) {
			CheckErr(invalidMap)
		}
	}
	if len(imported.mapdata) != imported.height {
		CheckErr(invalidMap)
	}
	//Check validity of file.

	return &imported

}

func (nm *nonomap) showMap() {

}

func (nm *nonomap) oneGame() {
	timer := NewPlaytime()
	go timer.showTime()
	estimatedTime := timer.timeResult()
	fmt.Println(estimatedTime)
}

func KeyStroke() {

}

func checkMark() {
}
