package customPath

import (
	"errors"
	"github.com/simp7/nonograminGo/file"
	"github.com/simp7/nonograminGo/file/localStorage"
	"os"
	"path"
)

type customPath struct {
	root string
	leaf []string
}

var (
	homePathErr = errors.New("HOME does not exist")
)

func new(leaf ...string) (customPath, error) {
	root, err := rootDir()
	return customPath{root, leaf}, err
}

func Get(name localStorage.PathName) (file.Path, error) {
	return new(string(name))
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
