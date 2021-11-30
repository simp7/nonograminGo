package main

import (
	"github.com/simp7/nonogram"
	"github.com/simp7/nonogram/db"
	"github.com/simp7/nonogram/db/formatter"
	"github.com/simp7/nonogram/db/local"
	"github.com/simp7/nonogram/unit/standard"
)

func main() {

	mapPrototype := standard.Prototype()

	maps := local.Map(mapPrototype.GetFormatter())
	settings := local.Setting(formatter.Json())
	languages := local.Language(formatter.Json())

	core := nonogram.New(db.New(maps, settings, languages), mapPrototype)

	rd := Controller(core)
	rd.Start()

}
