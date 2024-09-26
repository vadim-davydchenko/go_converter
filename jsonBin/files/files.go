package files

import (
	"os"
	"path/filepath"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func IsJSONFile(filePath string) bool {
	return filepath.Ext(filePath) == ".json"
}
