package formatter

import (
	"github.com/simp7/nonograminGo/errs"
	"github.com/simp7/nonograminGo/nonogram"
	"strconv"
	"strings"
)

type mapFormatter struct {
	data nonogram.Map
	raw  []byte
}

//TODO: Separate prototype from initializer(Maybe by using DI)

func Map(prototype nonogram.Map) *mapFormatter {
	formatter := new(mapFormatter)
	formatter.data = prototype
	formatter.raw = make([]byte, 0)
	return formatter
}

func (m *mapFormatter) Encode(i interface{}) error {

	switch i.(type) {
	case nonogram.Map:
		m.data = i.(nonogram.Map)
		return nil
	default:
		return errs.InvalidType
	}

}

func (m *mapFormatter) Decode(i interface{}) error {

	switch i.(type) {
	case *nonogram.Map:
		origin := i.(*nonogram.Map)
		*origin = m.data
		return nil
	default:
		return errs.InvalidType
	}

}

func (m *mapFormatter) GetRaw(content []byte) {

	m.raw = content

	data := string(content)
	builder := m.data.Builder()

	data = strings.TrimSpace(data)
	elements := strings.Split(data, "/")

	width, err := strconv.Atoi(elements[0])
	errs.Check(err)
	height, err := strconv.Atoi(elements[1])
	errs.Check(err)

	m.data = builder.BuildWidth(width).BuildHeight(height).BuildMap(elements[2:]).GetMap()
	m.data.CheckValidity()

}

func (m *mapFormatter) Content() []byte {
	return m.raw
}
