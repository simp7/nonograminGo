package localStorage

import (
	"github.com/simp7/nonograminGo/file"
	"os"
)

func MkDir(path file.Path) error {
	return os.Mkdir(path.String(), 0755)
}

func ReadFile(path file.Path) ([]byte, error) {
	return os.ReadFile(path.String())
}

func ReadDir(path file.Path) ([]os.DirEntry, error) {
	return os.ReadDir(path.String())
}

func WriteFile(path file.Path, data []byte) error {
	return os.WriteFile(path.String(), data, 0644)
}

func IsThere(path file.Path) bool {

	_, err := ReadDir(path)
	if err == nil {
		return true
	}
	_, err = ReadFile(path)
	return err == nil

}
