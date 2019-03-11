package util

import (
	"errors"
	"log"
)

var (
	invalidMap = errors.New("Map file has been broken.")
)

func CheckErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
