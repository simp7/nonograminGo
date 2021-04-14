package customPath

import (
	"errors"
	"github.com/simp7/nonograminGo/errs"
	"github.com/simp7/nonograminGo/nonogram/file"
	"os"
	"path"
)

type customPath struct {
	root string
	leaf []string
}

var (
	Root         = new("")
	SettingFile  = new("setting.json")
	MapsDir      = new("maps")
	LanguageDir  = new("language")
	LanguageFile = func(of string) file.Path { return LanguageDir.Append(of + ".json") }
	MapFile      = func(of string) file.Path { return MapsDir.Append(of) }
)

func new(leaf ...string) customPath {
	return customPath{rootDir(), leaf}
}

func (p customPath) String() string {
	result := p.root
	for _, v := range p.leaf {
		result = path.Join(result, v)
	}
	return result
}

func (p customPath) Append(newLeaf ...string) file.Path {
	return customPath{p.root, append(p.leaf, newLeaf...)}
}

func homeEnv() string {
	root, ok := os.LookupEnv("HOME")
	if !ok {
		errs.Check(errors.New("HOME does not exist"))
	}
	return root
}

func rootDir() string {
	return path.Join(homeEnv(), "nonogram")
}
