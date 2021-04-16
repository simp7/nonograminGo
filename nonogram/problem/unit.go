package problem

type unit struct {
	data []string
	max  int
}

func newUnit(data []string, max int) unit {
	return unit{data, max}
}

func (u unit) Get() []string {
	return u.data
}

func (u unit) Max() int {
	return u.max
}
