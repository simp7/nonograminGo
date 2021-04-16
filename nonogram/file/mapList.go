package file

type MapList interface {
	Current() []string
	Next()
	Prev()
	GetOrder() string
	GetMapName(from int) (name string, ok bool)
	GetCachedMapName() string
	Refresh()
}
