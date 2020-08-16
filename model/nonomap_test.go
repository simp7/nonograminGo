package model

import (
	"strings"
	"testing"
)

func TestNonomap_ShowBitMap(t *testing.T) {

	ans := [][]string{{"101", "010"}, {"111", "101", "010"}, {"1101", "0111", "1110", "1011"}}

	for i, m := range getExampleMap() {
		s := m.ShowBitMap()
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
	for i, s := range s1 {
		if strings.Compare(s, s2[i]) != 0 {
			t.Errorf("error in example %d -- expected : %s, actual : %s", idx, s2[i], s)
		}
	}
}

func getExampleMap() []*Nonomap {
	return []*Nonomap{NewNonomap("3/2/5/2"), NewNonomap("3/3/7/5/2"), NewNonomap("4/4/13/7/14/11")}
}
