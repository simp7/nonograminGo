package standard

import (
	"errors"
	"fmt"
	"github.com/simp7/nonograminGo/nonogram"
	"math"
	"strconv"
	"strings"
)

var (
	invalidType = errors.New("this file is not valid for map")
	invalidMap  = errors.New("map file has been broken")
)

type formatter struct {
	data nonogram.Map
	raw  []byte
}

func newFormatter() *formatter {
	f := new(formatter)
	f.data = Prototype()
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
		return invalidType
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
		return invalidType
	}

}

func (f *formatter) GetRaw(content []byte) error {

	f.raw = content

	data := string(content)

	data = strings.TrimSpace(data)
	elements := strings.Split(data, "/")

	width, err := strconv.Atoi(elements[0])
	if err != nil {
		return err
	}

	height, err := strconv.Atoi(elements[1])
	if err != nil {
		return err
	}

	bitmap, err := convertToBitmap(width, height, elements[2:])
	if err != nil {
		return err
	}

	f.data = f.data.CopyWithBitmap(bitmap)
	return f.data.CheckValidity()

}

func convertToBitmap(width, height int, data []string) ([][]bool, error) {

	result := make([][]bool, height)

	for i, v := range data {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, invalidMap
		}
		result[i] = getBitmapRow(num, width)
	}

	return result, nil

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
