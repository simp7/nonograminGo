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

func IsThere(path customPath) bool {

	_, err := ReadDir(path)
	if err == nil {
		return true
	}
	_, err = ReadFile(path)
	return err == nil

}
