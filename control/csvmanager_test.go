package control

import(
	"strings"
	"testing"
)

func TestNewCSVManager(t *testing.T) {
	cm := getCSVManagers()
	ans := []map[string]string{{"hello": "world"},{"Windows":"10","Linux":"Ubuntu","MacOS":"Catalina"}, {"color": "blue","number":"3","etc":"","name":"Park"}}
	for i, n := range cm {
		n.read()
		compareMap(i, n.csvMap, ans[i], t)
	}
}

func compareMap(idx int, m1 map[string]string, m2 map[string]string, t *testing.T) {
	for i, s := range m2 {
		if strings.Compare(s, m1[i]) != 0 {
			t.Errorf("error in example %d -- expected : %s, actual : %s", idx, s, m1[i])
		}
	}
}

func getCSVManagers() []*CSVManager{
	return []*CSVManager{NewCSVManager("test1"), NewCSVManager("test2"), NewCSVManager("test3")}
}