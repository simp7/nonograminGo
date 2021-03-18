package fileFormatter

import (
	"github.com/simp7/nonograminGo/nonogram"
	"github.com/simp7/nonograminGo/nonogram/nonomap"
	"github.com/simp7/nonograminGo/util"
	"reflect"
	"strconv"
	"strings"
)

type mapFormatter struct {
	data nonogram.Map
	raw  []byte
}

func Map() nonogram.FileFormatter {
	formatter := new(mapFormatter)
	formatter.raw = make([]byte, 0)
	return formatter
}

func (m *mapFormatter) Encode(i interface{}) error {
	switch i.(type) {
	case nonogram.Map:
		m.data = i.(nonogram.Map)
	default:
		return util.InvalidType
	}
	return nil
}

func (m *mapFormatter) Decode(i interface{}) error {

	rv := reflect.ValueOf(i)
	switch rv.Type() {
	default:
		return util.InvalidType
	case reflect.TypeOf(m.data):
		rv.Elem().Set(reflect.ValueOf(m.data).Elem())
	}
	return nil

}

func (m *mapFormatter) GetRaw(content []byte) {

	m.raw = content

	data := string(content)
	builder := nonomap.NewNonomapBuilder()

	data = strings.TrimSpace(data)
	elements := strings.Split(data, "/")

	width, err := strconv.Atoi(elements[0])
	util.CheckErr(err)
	height, err := strconv.Atoi(elements[1])
	util.CheckErr(err)

	m.data = builder.BuildWidth(width).BuildHeight(height).BuildMap(elements[2:]).GetMap()
	m.data.CheckValidity()

}

func (m *mapFormatter) Content() []byte {
	return m.raw
}
