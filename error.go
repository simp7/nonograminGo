package main

import (
	"errors"
	"log"
)

var (
	timeBelowZero = errors.New("The value of time is less than 0.")
)

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
