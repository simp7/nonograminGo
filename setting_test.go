package main

import (
	"fmt"
	"github.com/simp7/nonograminGo/asset"
	"strings"
	"testing"
)

func TestGetSetting(t *testing.T) {
	s := asset.GetSetting()
	fmt.Println(s.Text)
	a := [][]string{s.MainMenu(), s.GetHelp(), {s.RequestMapName(), s.FileNotExist()}}
	b := [][]string{{"----------", " NONOGRAM", "----------", "", "Press number you want to select.", "", "1. START", "2. CREATE", "3. HELP", "4. CREDIT", "5. EXIT"}, {"    MANUAL", "--------------", "Arrow Key : Move cursor", "Space or Z : Fill the cell", "X : Check the cell that is supposed to be blank", "Enter(create mode) : Save the map that player creates", "Esc : Get out of current game/display"}, {"Write map name that you want to create", "File doesn't exist."}}
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
