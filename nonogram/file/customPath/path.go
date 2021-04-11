package customPath

import (
	_ "embed"
	"errors"
	"github.com/simp7/nonograminGo/errs"
	"github.com/simp7/nonograminGo/nonogram/file"
	"os"
	"path"
)

//go embed:skel

type customPath struct {
	root string
	leaf []string
}

var (
	DefaultSettingFile = source("default_setting.json")
	SettingFile        = real("setting.json")
	DefaultMapsDir     = source("default_maps")
	MapsDir            = real("maps")
	DefaultLanguageDir = source("language")
	LanguageDir        = real("language")
	LanguageFile       = func(of string) file.Path { return LanguageDir.Append(of + ".json") }
	MapFile            = func(of string) file.Path { return MapsDir.Append(of) }
)

func newPath(root string, leaf ...string) file.Path {
	return customPath{root, leaf}
}

func real(leaf ...string) file.Path {
	return newPath(rootDir(), leaf...)
}

func source(leaf ...string) file.Path {
	return newPath("skel", leaf...)
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
