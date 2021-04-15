package file

import "github.com/simp7/nonograminGo/nonogram"

type MapList interface {
	Current() []string
	Next()
	Prev()
	GetOrder() string
	GetMapName(from int) (name string, ok bool)
	GetCachedMapName() string
	CreateMap(mapData nonogram.Map, name string)
}
