package model

import (
	"github.com/simp7/nonograminGo/util"
	"strconv"
)

type NonomapBuilder interface {
	BuildHeight(int) NonomapBuilder
	BuildWidth(int) NonomapBuilder
	BuildMap([]string) NonomapBuilder
	GetMap() Nonomap
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
		util.CheckErr(err)
	}

	return b
}

func (b *nonomapBuilder) GetMap() Nonomap {
	return b.data
}
