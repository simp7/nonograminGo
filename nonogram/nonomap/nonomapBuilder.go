package nonomap

import (
	"github.com/simp7/nonograminGo/errs"
	"github.com/simp7/nonograminGo/nonogram"
	"strconv"
)

type builder struct {
	data *nonomap
}

func NewBuilder() nonogram.MapBuilder {
	b := new(builder)
	b.data = new(nonomap)
	return b
}

func (b *builder) Height(h int) nonogram.MapBuilder {
	b.data.Height = h
	return b
}

func (b *builder) Width(w int) nonogram.MapBuilder {
	b.data.Width = w
	return b
}

func (b *builder) Map(content []string) nonogram.MapBuilder {

	for _, v := range content {
		tmp, err := strconv.Atoi(v)
		b.data.MapData = append(b.data.MapData, tmp)
		errs.Check(err)
	}

	b.bitMap()

	return b

}

func (b *builder) bitMap() {

	nmap := b.data

	b.data.Bitmap = make([][]bool, nmap.Height)
	for n := range nmap.Bitmap {
		b.data.Bitmap[n] = make([]bool, nmap.Width)
	}

	for i, v := range nmap.MapData {
		b.bitMapByRow(i, v)
	}

}

func (b *builder) bitMapByRow(y, rowValue int) {

	width := b.data.Width
	v := rowValue

	for i := 1; i <= width; i++ {
		b.data.Bitmap[y][width-i] = v%2 == 1
		v = v / 2
	}

}

func (b *builder) Build() nonogram.Map {
	return b.data
}
