package asset

import "sync"

type Setting interface {
}

type setting struct {
	color Color
	text  Text
}

var instance Setting
var once sync.Once

func GetInstance() Setting {
	once.Do(func() {
		s := new(setting)
		s.color = new(color)
	})
	return instance
}
