package customPath

import (
	"errors"
	"github.com/simp7/nonograminGo/errs"
	"github.com/simp7/nonograminGo/nonogram"
	"os"
	"path"
)

type customPath struct {
	root string
	leaf []string
}

func newPath(root string, leaf ...string) nonogram.Path {
	return customPath{root, leaf}
}

func Real(leaf ...string) nonogram.Path {
	return newPath(homePath(), leaf...)
}

func Source(leaf ...string) nonogram.Path {
	return newPath("files", leaf...)
}

func (p customPath) String() string {
	result := p.root
	for _, v := range p.leaf {
		result = path.Join(result, v)
	}
	return result
}

func (p customPath) Append(newLeaf ...string) nonogram.Path {
	return customPath{p.root, append(p.leaf, newLeaf...)}
}

func homePath() string {
	root, ok := os.LookupEnv("HOME")
	if !ok {
		errs.Check(errors.New("HOME does not exist"))
	}
	return root
}
