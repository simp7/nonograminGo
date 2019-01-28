package main

import (
	"fmt"
	"time"
)

//This file deals with timer.
type Playtime struct {
	ticker    time.Ticker
	startTime time.Time
	clock     chan int
	stop      chan struct{}
}

func NewPlaytime() *Playtime {
	var p Playtime
	p.ticker = *time.NewTicker(time.Second)
	p.startTime = time.Now()
	p.clock = make(chan int)
	p.stop = make(chan struct{})
	return &p
}

func (p *Playtime) timePassed() { //would be goroutine
	present := 0

	select {
	case <-p.ticker.C:
		present += 1
		p.clock <- present
	case <-p.stop:
		return
	}

}

func (p *Playtime) timeResult() float64 {
	result := time.Now().Sub(p.startTime)
	if result < 0 {
		CheckErr(timeBelowZero)
	}
	close(p.stop)
	close(p.clock)
	return result.Seconds()
}

func (p *Playtime) showTime() {
	fmt.Println(<-p.clock)
}
