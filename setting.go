package main

import (
	"github.com/simp7/nonogram/setting"
)

type Config struct {
	Color
	Text
	Language   string
	NameMax    int
	DefaultPos Pos
}

func AdjustConfig(conf setting.Config) Config {
	return Config{Light(), AdaptText(conf.Text), conf.Language, 30, Pos{5, 5}}
}
