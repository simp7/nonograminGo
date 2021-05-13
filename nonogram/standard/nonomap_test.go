package standard

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNonomap_CreateProblem(t *testing.T) {

	nm := Map()
	copied := nm.CopyWithBitmap([][]bool{{true, false, true}, {false, true, false}}).CreateProblem()
	assert.Equal(t, newProblem([][]int{{1, 1}, {1}}, [][]int{{1}, {1}, {1}}, 2, 1), copied)
	table := []struct {
		bitmap     [][]bool
		horizontal [][]int
		vertical   [][]int
		hMax       int
		vMax       int
		msg        string
	}{
		{[][]bool{{true}}, [][]int{{1}}, [][]int{{1}}, 1, 1, "one cell"},
		{[][]bool{{true, false, true}, {false, true, false}}, [][]int{{1, 1}, {1}}, [][]int{{1}, {1}, {1}}, 2, 1, "problem that consists with 1."},
		{[][]bool{{true, false, true}, {true, true, true}}, [][]int{{1, 1}, {3}}, [][]int{{2}, {1}, {2}}, 2, 1, "many cell test."},
	}

	for _, v := range table {
		wanted := newProblem(v.horizontal, v.vertical, v.hMax, v.vMax)
		got := nm.CopyWithBitmap(v.bitmap).CreateProblem()
		assert.Equal(t, wanted, got, v.msg)
	}

}
