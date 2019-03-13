package util

import (
	"errors"
	"log"
)

var (
	InvalidMap = errors.New("Map file has been broken.")
)

func CheckErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
