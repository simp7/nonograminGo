package localStorage

import (
	"errors"
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

func pathTo(leaf ...string) (customPath, error) {
	root, err := rootDir()
	return customPath{root, leaf}, err
}

func Get(name PathName) (customPath, error) {
	return pathTo(string(name))
}

func (p customPath) String() string {
	result := p.root
	for _, v := range p.leaf {
		result = path.Join(result, v)
	}
	return result
}

func (p customPath) Append(newLeaf ...string) customPath {
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
