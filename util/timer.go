package util

import (
	"fmt"
	"time"
)

/*
This file deals with timer.
Playtime can be used in show and record playtime of current map.
*/

type Playtime struct {
	ticker time.Ticker
	Clock  chan string
	Stop   chan struct{}
}

func NewPlaytime() *Playtime {

	var p Playtime

	p.ticker = *time.NewTicker(time.Second)
	p.Clock = make(chan string)
	p.Stop = make(chan struct{})
	go p.timePassed()

	return &p

}

/*
This function will send the seconds that has passed during gameplay.
This function will be called in NewPlaytime.
This function should be called in goroutine.
*/

func (p *Playtime) timePassed() {

	present := 0
	p.Clock <- convertTimeFormat(present)

	for {
		select {

		case <-p.ticker.C:
			present += 1
			p.Clock <- convertTimeFormat(present)

		case <-p.Stop:
			p.Clock <- convertTimeFormat(present) //To prevent situation that p.Clock channel is empty.
			p.ticker.Stop()
			close(p.Clock)
			return

		}
	}

}

/*
This function returns seconds that has passsed during gameplay.
This function will be called when player finished the map.
*/

func (p *Playtime) TimeResult() string {

	close(p.Stop)
	return <-p.Clock

}

/*
This function will be called when player ends the game without solving.
*/

func (p *Playtime) EndWithoutResult() {

	close(p.Stop)

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
