package customPath

import (
	"errors"
	"github.com/simp7/nonograminGo/errs"
	"github.com/simp7/nonograminGo/file"
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
	homePathErr  = errors.New("HOME does not exist")
)

func new(leaf ...string) customPath {
	root, err := rootDir()
	errs.Check(err)
	return customPath{root, leaf}
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

func homeEnv() (string, error) {
	root, ok := os.LookupEnv("HOME")
	if !ok {
		return "", homePathErr
	}
	return root, nil
}

func rootDir() (string, error) {
	home, err := homeEnv()
	if err != nil {
		return "", err
	}
	return path.Join(home, "nonogram"), nil
}
