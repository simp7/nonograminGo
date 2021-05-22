package client

//Direction represents directions for playing nonogram.
type Direction uint8

const (
	Up Direction = iota
	Down
	Left
	Right
)
