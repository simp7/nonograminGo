package model

import (
	"github.com/simp7/nonograminGo/asset"
	"github.com/simp7/nonograminGo/util"
	"math"
	"strconv"
	"strings"
)

type mapFormatter struct {
	data *Nonomap
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
		m.data = i.(*Nonomap)
	default:
		return util.InvalidType
	}
	return nil
}

func (m *mapFormatter) Decode(i interface{}) error {
	switch i.(type) {
	case *Nonomap:
		i = m.data
	default:
		return util.InvalidType
	}
	return nil
}

func (m *mapFormatter) GetRaw(from []byte) {
	var err error
	m.raw = from
	data := string(from)
	imported := m.data
	setting := asset.GetSetting()

	data = strings.TrimSpace(data)
	elements := strings.Split(data, "/")
	//Extract all data from wanted file.

	imported.Width, err = strconv.Atoi(elements[0])
	imported.Height, err = strconv.Atoi(elements[1])
	util.CheckErr(err)
	//Extract map's size from file.

	for _, v := range elements[2:] {
		temp, err := strconv.Atoi(v)
		imported.MapData = append(imported.MapData, temp)
		util.CheckErr(err)
	}
	//Extract map's answer from file.

	if imported.Height > setting.HeightMax || imported.Width > setting.WidthMax || imported.Height <= 0 || imported.Width <= 0 {
		util.CheckErr(util.InvalidMap)
	} //Check if Height and Width meets criteria of size.

	for _, v := range imported.MapData {
		if float64(v) >= math.Pow(2, float64(imported.Width)) {
			util.CheckErr(util.InvalidMap)
		} //Check whether Height matches MapData.
	}
	if len(imported.MapData) != imported.Height {
		util.CheckErr(util.InvalidMap)
	} //Check whether Height matches MapData.

	//Check validity of file.
	imported.Bitmap = convertToBitmap(imported.Width, imported.Height, imported.MapData)
}

func (m *mapFormatter) Content() []byte {
	return m.raw
}
