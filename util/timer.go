package util

import (
	"strconv"
	"time"
)

/*
This file deals with timer.
Playtime can be used in show and record playtime of current map.
*/

type Playtime struct {
	ticker time.Ticker
	Clock  chan string
	stop   chan struct{}
}

func NewPlaytime() *Playtime {

	var p Playtime

	p.ticker = *time.NewTicker(time.Second)
	p.Clock = make(chan string)
	p.stop = make(chan struct{})

	return &p

}

/*
This function will send the seconds that has passed during gameplay.
This function will be called when the game starts.
This function should be called in goroutine.
*/

func (p *Playtime) TimePassed() {

	present := 0

	for {
		select {

		case <-p.ticker.C:
			present += 1
			p.Clock <- strconv.Itoa(present)

		case <-p.stop:
			p.Clock <- strconv.Itoa(present) //To prevent situation that p.clock channel is empty.
			p.ticker.Stop()
			return

		}
	}

}

/*
This function returns seconds that has passsed during gameplay.
This function will be called when player finished the map.
*/

func (p *Playtime) TimeResult() string {

	close(p.stop)
	defer close(p.Clock)

	return <-p.Clock

}
