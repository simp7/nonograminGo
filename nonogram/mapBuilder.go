package nonogram

type MapBuilder interface {
	Height(int) MapBuilder
	Width(int) MapBuilder
	Map([]string) MapBuilder
	Build() Map
}
