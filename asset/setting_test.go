package asset

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetSetting(t *testing.T) {
	s := GetSetting()
	fmt.Println(s.Text)
	a := [][]string{s.MainMenu(), s.GetHelp(), {s.RequestMapName(), s.FileNotExist()}}
	b := [][]string{StringMainMenu, StringHelp, {StringHeaderMapName, StringMsgFileNotExist}}
	for i := range a {
		compareTexts(a[i], b[i], t)
	}
}

func compareTexts(a []string, b []string, t *testing.T) {
	for i := range a {
		compareText(a[i], b[i], t)
	}
}

func compareText(a string, b string, t *testing.T) {
	if strings.Compare(a, b) != 0 {
		t.Errorf("Result should be %s, got %s", b, a)
	}
}
