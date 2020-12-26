package asset

import "sync"

type Setting interface {
}

type setting struct {
	color Color
}

var instance Setting
var once sync.Once

func getInstance() Setting {
	once.Do(func() {
		s := new(setting)
		s.color = new(color)
	})
	return instance
}
