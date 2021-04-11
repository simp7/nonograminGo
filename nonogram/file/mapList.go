package file

type MapList interface {
	GetAll() []string
	Next()
	Prev()
	GetOrder() string
	GetMapName(from int) (name string, ok bool)
	GetCachedMapName() string
	CreateMap(name string, width int, height int, bitmap [][]bool)
	Refresh()
}
