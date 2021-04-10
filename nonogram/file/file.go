package file

import (
	"os"
)

func MkDir(path Path) error {
	return os.Mkdir(path.String(), 0755)
}

func ReadFile(path Path) ([]byte, error) {
	return os.ReadFile(path.String())
}

func ReadDir(path Path) ([]os.DirEntry, error) {
	return os.ReadDir(path.String())
}

func WriteFile(path Path, data []byte) error {
	return os.WriteFile(path.String(), data, 0644)
}
