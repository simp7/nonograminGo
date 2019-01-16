package main

import (
	"time"
)

type Playtime struct {
	ticker    time.Ticker
	startTime time.Time
	clock     chan int
}

func NewPlaytime() *Playtime {
	var p Playtime
	p.ticker = *time.NewTicker(time.Second)
	p.startTime = time.Now()
	p.clock = make(chan int)
	return &p
}

func (p *Playtime) timePassed() { //would be go routine
	present := 0

	select {
	case <-p.ticker.C:
		present += 1
		p.clock <- present
	}

}

func (p *Playtime) timeResult() float64 {
	result := time.Now().Sub(p.startTime)
	return result.Seconds()
}
