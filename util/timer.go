package util

import (
	"fmt"
	"time"
)

/*
This file deals with timer.
Timer can be used in show and record playtime of current map.
*/

type Timer interface {
	GetResult() string
	End()
	Do(func(current string))
}

type timer struct {
	time.Ticker
	clock   chan string
	stopper chan struct{}
}

func NewPlaytime() Timer {

	var p timer

	p.Ticker = *time.NewTicker(time.Second)
	p.clock = make(chan string)
	p.stopper = make(chan struct{})
	go p.timePassed()

	return &p

}

/*
This function will send the seconds that has passed during game.
This function will be called in NewPlaytime.
This function should be called in goroutine.
*/

func (p *timer) timePassed() {

	present := 0
	p.clock <- convertTimeFormat(present)

	for {
		select {

		case <-p.C:
			present += 1
			p.clock <- convertTimeFormat(present)

		case <-p.stopper:
			p.clock <- convertTimeFormat(present) //To prevent situation that p.clock channel is empty.
			p.Stop()
			close(p.clock)
			return

		}
	}

}

/*
This function returns seconds that has passed during game.
This function will be called when player finished the map.
*/

func (p *timer) GetResult() string {

	p.End()
	return <-p.clock

}

/*
This function will be called when player ends the game without solving.
*/

func (p *timer) End() {

	close(p.stopper)

}

func (p *timer) Do(someFunc func(current string)) {
	for {
		select {
		case current := <-p.clock:
			someFunc(current)
		case <-p.stopper:
			return
		}
	}
}

func convertTimeFormat(totalTime int) string {

	minutes := totalTime / 60
	seconds := totalTime % 60

	if seconds < 10 {
		return fmt.Sprintf("%d:0%d", minutes, seconds)
	} else {
		return fmt.Sprintf("%d:%d", minutes, seconds)
	}

}
