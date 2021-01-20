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

func StartTimer() Timer {

	t := new(timer)

	t.Ticker = *time.NewTicker(time.Second)
	t.clock = make(chan string)
	t.stopper = make(chan struct{})
	go t.timePassed()

	return t

}

/*
This function will send the seconds that has passed during game.
This function will be called in StartTimer.
This function should be called in goroutine.
*/

func (t *timer) timePassed() {

	present := 0
	t.clock <- convertTimeFormat(present)

	for {
		select {

		case <-t.C:
			present += 1
			t.clock <- convertTimeFormat(present)
			if present > 3600*24-1 {
				t.Stop()
			}

		case <-t.stopper:
			t.clock <- convertTimeFormat(present) //To prevent situation that t.clock channel is empty.
			t.Stop()
			close(t.clock)
			return

		}
	}

}

/*
This function returns seconds that has passed during game.
This function will be called when player finished the map.
*/

func (t *timer) GetResult() string {

	t.End()
	return <-t.clock

}

/*
This function will be called when player ends the game without solving.
*/

func (t *timer) End() {

	close(t.stopper)

}

func (t *timer) Do(someFunc func(current string)) {
	for {
		select {
		case current := <-t.clock:
			someFunc(current)
		case <-t.stopper:
			return
		}
	}
}

func convertTimeFormat(totalTime int) string {

	seconds := totalTime % 60
	minutes := (totalTime % 3600) / 60
	hours := totalTime / 3600

	if seconds < 10 {
		return fmt.Sprintf("%d:0%d", minutes, seconds)
	} else if hours == 0 {
		return fmt.Sprintf("%d:%d", minutes, seconds)
	} else {
		return fmt.Sprintf("%d:%d:%d", hours, minutes, seconds)
	}

}
