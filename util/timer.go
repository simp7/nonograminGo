package util

import (
	"time"
)

/*
This file deals with timer.
Playtime can be used in show and record playtime of current map.
*/

type Playtime struct {
	ticker time.Ticker
	clock  chan int
	stop   chan struct{}
}

func NewPlaytime() *Playtime {

	var p Playtime

	p.ticker = *time.NewTicker(time.Second)
	p.clock = make(chan int)
	p.stop = make(chan struct{})

	return &p

}

/*
This function will send the seconds that has passed during gameplay.
This function will be called when the game starts.
This function should be called in goroutine.
*/

func (p *Playtime) timePassed() {

	present := 0

	for {
		select {

		case <-p.ticker.C:
			present += 1
			p.clock <- present

		case <-p.stop:
			p.clock <- present //To prevent situation that p.clock channel is empty.
			return

		}
	}

}

/*
This function returns seconds that has passsed during gameplay.
This function will be called when player finished the map.
*/

func (p *Playtime) timeResult() int {

	close(p.stop)
	close(p.clock)

	return <-p.clock

}
