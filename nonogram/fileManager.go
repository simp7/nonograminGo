package nonogram

type FileManager interface {
	GetMapList() []string
	NextList()
	PrevList()
	GetOrder() string
	GetMapDataByNumber(int) (nonomap Map, ok bool)
	GetMapDataByName(string) (nonomap Map, ok bool)
	GetCurrentMapName() string
	CreateMap(name string, width int, height int, bitmap [][]bool)
	RefreshMapList()
}
