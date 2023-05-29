package utils

import (
	"os"
)

func CreateTempDir() (string, error) {
	tempDir, err := os.MkdirTemp("", "tmpGithub")
	if err != nil {
		return "", err
	}

	return tempDir, nil
}

func CleanupTempDir(source string) error {
	err := os.RemoveAll(source)
	if err != nil {
		return err
	}
	return nil
}
