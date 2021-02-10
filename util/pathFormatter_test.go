package util

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestPathFormatter_GetPath(t *testing.T) {
	f := GetPathFormatter()
	sep := string(filepath.Separator)

	ans := []string{".", fmt.Sprintf("hello"), fmt.Sprintf("main%ssomeFile", sep)}
	input := []string{f.GetPath(), f.GetPath("hello"), f.GetPath("main", "someFile")}

	for i, v := range ans {
		compare(i, v, input[i], t)
	}
}

func compare(idx int, wanted, get string, t *testing.T) {
	if wanted != get {
		t.Errorf("Errors in %d : wanted %s, but got %s", idx, wanted, get)
	} else {
		t.Logf("Example %d has been passed", idx)
	}
}
