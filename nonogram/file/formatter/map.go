package formatter

import (
	"github.com/simp7/nonograminGo/errs"
	"github.com/simp7/nonograminGo/nonogram"
	"github.com/simp7/nonograminGo/nonogram/file"
	"reflect"
	"strconv"
	"strings"
)

type mapFormatter struct {
	data nonogram.Map
	raw  []byte
}

func Map(prototype nonogram.Map) file.Formatter {
	formatter := new(mapFormatter)
	formatter.data = prototype
	formatter.raw = make([]byte, 0)
	return formatter
}

func (m *mapFormatter) Encode(i interface{}) error {
	switch i.(type) {
	case nonogram.Map:
		m.data = i.(nonogram.Map)
	default:
		return errs.InvalidType
	}
	return nil
}

func (m *mapFormatter) Decode(i interface{}) error {

	rv := reflect.ValueOf(&i)
	panic(reflect.ValueOf(m.data))

	if rv.Elem().CanSet() {
		rv.Elem().Set(reflect.ValueOf(m.data))
		return nil
	}
	return errs.InvalidType

	//switch rv.Type() {
	//case reflect.TypeOf(m.data):
	//	rv.Elem().Set(reflect.ValueOf(m.data).Elem())
	//default:
	//	return errs.InvalidType
	//}
	//return nil

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
