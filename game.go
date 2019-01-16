package main

import (
	"fmt"
	"io/ioutil"
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
	checkErr(err)

	elements := strings.Split(string(file), "/")

	imported.width, err = strconv.Atoi(elements[0])
	imported.height, err = strconv.Atoi(elements[1])
	checkErr(err)

	for _, v := range elements[2:] {
		temp, err := strconv.Atoi(v)
		imported.mapdata = append(imported.mapdata, temp)
		checkErr(err)
	}

	return &imported

}

func (nm *nonomap) oneGame() {
	timer := NewPlaytime()
	go timer.showTime()
	go check()
	estimatedTime := timer.timeResult()
	fmt.Println(estimatedTime)
}

func check() {
}
