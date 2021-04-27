package localStorage

import (
	"os"
)

func MkDir(path customPath) error {
	return os.Mkdir(path.String(), 0755)
}

func ReadFile(path customPath) ([]byte, error) {
	return os.ReadFile(path.String())
}

func ReadDir(path customPath) ([]os.DirEntry, error) {
	return os.ReadDir(path.String())
}

func WriteFile(path customPath, data []byte) error {
	return os.WriteFile(path.String(), data, 0644)
}

func IsInitial() bool {

	path, _ := Get(ROOT)

	_, err1 := ReadDir(path)
	_, err2 := ReadFile(path)

	return err1 != nil && err2 != nil

}
