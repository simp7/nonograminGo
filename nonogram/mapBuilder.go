package nonogram

type MapBuilder interface {
	BuildHeight(int) MapBuilder
	BuildWidth(int) MapBuilder
	BuildMap([]string) MapBuilder
	GetMap() Map
}
