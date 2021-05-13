package standard

type unit struct {
	data [][]int
	max  int
}

func newUnit(data [][]int, max int) unit {
	return unit{data, max}
}

func (u unit) Get(idx int) []int {
	return u.data[idx]
}

func (u unit) Max() int {
	return u.max
}
