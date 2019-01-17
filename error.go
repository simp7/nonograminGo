package main

import (
	"errors"
	"log"
)

var (
	timeBelowZero = errors.New("The value of time is less than 0.")
	invalidMap    = errors.New("Map file has been broken.")
)

func CheckErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
