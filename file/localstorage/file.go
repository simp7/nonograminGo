package localstorage

import (
	"os"
)

func mkDir(path customPath) error {
	return os.Mkdir(path.String(), 0755)
}

func readFile(path customPath) ([]byte, error) {
	return os.ReadFile(path.String())
}

func readDir(path customPath) ([]os.DirEntry, error) {
	return os.ReadDir(path.String())
}

func writeFile(path customPath, data []byte) error {
	return os.WriteFile(path.String(), data, 0644)
}

func isInitial() bool {

	path, _ := get(ROOT)

	_, err1 := readDir(path)
	_, err2 := readFile(path)

	return err1 != nil && err2 != nil

}
