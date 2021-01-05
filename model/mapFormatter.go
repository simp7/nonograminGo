package model

import (
	"fmt"
	"github.com/simp7/nonograminGo/util"
	"strconv"
	"strings"
)

type mapFormatter struct {
	data Nonomap
	raw  []byte
}

func NewMapFormatter() util.FileFormatter {
	formatter := new(mapFormatter)
	formatter.raw = make([]byte, 0)
	return formatter
}

func (m *mapFormatter) Encode(i interface{}) error {
	switch i.(type) {
	case Nonomap:
		m.data = i.(Nonomap)
	default:
		return util.InvalidType
	}
	return nil
}

func (m *mapFormatter) Decode(i interface{}) error {
	switch i.(type) {
	case Nonomap:
		i = m.data
	default:
		return util.InvalidType
	}
	return nil
}

func (m *mapFormatter) GetRaw(content []byte) {

	copy(m.raw, content)

	data := string(content)
	builder := NewNonomapBuilder()

	data = strings.TrimSpace(data)
	elements := strings.Split(data, "/")

	width, err := strconv.Atoi(elements[0])
	util.CheckErr(err)
	height, err := strconv.Atoi(elements[1])
	util.CheckErr(err)

	m.data = builder.BuildWidth(width).BuildHeight(height).BuildMap(elements[:2]).GetMap()
	m.data.checkValidity()
	fmt.Println(m.data)

}

func (m *mapFormatter) Content() []byte {
	return m.raw
}
