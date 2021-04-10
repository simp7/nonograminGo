package file

import "github.com/simp7/nonograminGo/nonogram"

type MapsLoader interface {
	GetMapList() []string
	NextList()
	PrevList()
	GetOrder() string
	GetMapDataByNumber(int) (nonomap nonogram.Map, ok bool)
	GetMapDataByName(string) (nonomap nonogram.Map, ok bool)
	GetCurrentMapName() string
	CreateMap(name string, width int, height int, bitmap [][]bool)
	RefreshMapList()
}
