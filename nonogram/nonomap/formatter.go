package nonomap

import (
	"fmt"
	"github.com/simp7/nonograminGo/errs"
	"github.com/simp7/nonograminGo/nonogram"
	"math"
	"strconv"
	"strings"
)

type formatter struct {
	data nonogram.Map
	raw  []byte
}

func Formatter() *formatter {
	f := new(formatter)
	f.data = New()
	f.raw = make([]byte, 0)
	return f
}

func (f *formatter) Encode(i interface{}) error {

	switch i.(type) {
	case *nonomap:
		f.data = i.(*nonomap)
		f.raw = convert(i.(*nonomap))
		return nil
	default:
		return errs.InvalidType
	}

}

func convert(nmap *nonomap) []byte {

	result := fmt.Sprintf("%d/%d", nmap.GetWidth(), nmap.GetHeight())

	for _, row := range nmap.Bitmap {
		result += fmt.Sprintf("/%d", getRowValue(nmap.GetWidth(), row))
	}

	return []byte(result)

}

func getRowValue(width int, row []bool) int {

	result := 0

	for i, v := range row {
		if v {
			result += int(math.Pow(2, float64(width-i-1)))
		}
	}

	return result

}

func (f *formatter) Decode(i interface{}) error {

	switch i.(type) {
	case *nonogram.Map:
		origin := i.(*nonogram.Map)
		*origin = f.data
		return nil
	default:
		return errs.InvalidType
	}

}

func (f *formatter) GetRaw(content []byte) {

	f.raw = content

	data := string(content)

	data = strings.TrimSpace(data)
	elements := strings.Split(data, "/")

	width, err := strconv.Atoi(elements[0])
	errs.Check(err)
	height, err := strconv.Atoi(elements[1])
	errs.Check(err)

	f.data = NewByBitMap(convertToBitmap(width, height, elements[2:]))
	errs.Check(f.data.CheckValidity())

}

func convertToBitmap(width, height int, data []string) (result [][]bool) {

	result = make([][]bool, height)

	for i, v := range data {
		num, err := strconv.Atoi(v)
		if err != nil {
			errs.Check(invalidMap)
		}
		result[i] = getBitmapRow(num, width)
	}

	return

}

func getBitmapRow(value, width int) []bool {

	result := make([]bool, width)

	for i := 1; i <= width; i++ {
		result[width-i] = value%2 == 1
		value = value / 2
	}

	return result

}

func (f *formatter) Content() []byte {
	return f.raw
}
