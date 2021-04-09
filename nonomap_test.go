package main

import (
	"github.com/simp7/nonograminGo/errs"
	"github.com/simp7/nonograminGo/nonogram"
	"github.com/simp7/nonograminGo/nonogram/fileFormatter"
	"github.com/simp7/nonograminGo/nonogram/nonomap"
	"testing"
)

func TestNonomap_ShowBitMap(t *testing.T) {

	ans := [][]string{{"101", "010"}, {"111", "101", "010"}, {"1101", "0111", "1110", "1011"}}

	for i, m := range getExampleMap() {
		s := m.BitmapToStrings()
		strArrayCompare(s, ans[i], i, t)
	}

}

func TestNonomap_ShowProblemHorizontal(t *testing.T) {
	ans := [][]string{{"11", "1"}, {"3", "11", "1"}, {"21", "3", "3", "12"}}

	for i, m := range getExampleMap() {
		s := m.ShowProblemHorizontal()
		strArrayCompare(s, ans[i], i, t)
	}
}

func TestNonomap_ShowProblemVertical(t *testing.T) {
	ans := [][]string{{"1", "1", "1"}, {"2", "11", "2"}, {"12", "3", "3", "21"}}

	for i, m := range getExampleMap() {
		s := m.ShowProblemVertical()
		strArrayCompare(s, ans[i], i, t)
	}
}

func strArrayCompare(s1 []string, s2 []string, idx int, t *testing.T) {
	for i := range s1 {
		stringCompare(s1[i], s2[i], idx, t)
	}
}

func stringCompare(s1 string, s2 string, idx int, t *testing.T) {
	if s1 != s2 {
		t.Errorf("error in example %d -- expected : %s, actual %s", idx, s2, s1)
	} else {
		t.Logf("example %d has been passed", idx)
	}
}

func getExampleMap() []nonogram.Map {

	newMap := func(data string) nonogram.Map {

		result := nonomap.New()
		f := fileFormatter.Map()

		f.GetRaw([]byte(data))
		errs.Check(f.Decode(result))

		return result

	}

	return []nonogram.Map{newMap("3/2/5/2"), newMap("3/3/7/5/2"), newMap("4/4/13/7/14/11")}

}
