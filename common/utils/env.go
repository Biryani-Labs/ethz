package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func ExpandPath(path string) (string, error) {
	if strings.HasPrefix(path, "~/") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		path = filepath.Join(homeDir, path[2:])
	}
	return filepath.Abs(path)
}
