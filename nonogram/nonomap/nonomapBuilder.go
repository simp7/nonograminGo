package nonomap

import (
	"github.com/simp7/nonograminGo/errs"
	"github.com/simp7/nonograminGo/nonogram"
	"strconv"
)

type NonomapBuilder interface {
	BuildHeight(int) NonomapBuilder
	BuildWidth(int) NonomapBuilder
	BuildMap([]string) NonomapBuilder
	GetMap() nonogram.Map
}

type nonomapBuilder struct {
	data *nonomap
}

func NewNonomapBuilder() NonomapBuilder {
	b := new(nonomapBuilder)
	b.data = new(nonomap)
	return b
}

func (b *nonomapBuilder) BuildHeight(h int) NonomapBuilder {
	b.data.Height = h
	return b
}

func (b *nonomapBuilder) BuildWidth(w int) NonomapBuilder {
	b.data.Width = w
	return b
}

func (b *nonomapBuilder) BuildMap(content []string) NonomapBuilder {

	for _, v := range content {
		tmp, err := strconv.Atoi(v)
		b.data.MapData = append(b.data.MapData, tmp)
		errs.Check(err)
	}

	b.buildBitMap()

	return b

}

func (b *nonomapBuilder) buildBitMap() {

	nmap := b.data

	b.data.Bitmap = make([][]bool, nmap.Height)
	for n := range nmap.Bitmap {
		b.data.Bitmap[n] = make([]bool, nmap.Width)
	}

	for i, v := range nmap.MapData {
		b.buildBitMapByRow(i, v)
	}

}

func (b *nonomapBuilder) buildBitMapByRow(y, rowValue int) {

	width := b.data.Width
	v := rowValue

	for i := 1; i <= width; i++ {
		b.data.Bitmap[y][width-i] = v%2 == 1
		v = v / 2
	}

}

func (b *nonomapBuilder) GetMap() nonogram.Map {
	return b.data
}
