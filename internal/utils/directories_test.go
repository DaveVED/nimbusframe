package utils

import (
	"os"
	"testing"
)

func TestCreateTempDir_Positive(t *testing.T) {
	dir, err := CreateTempDir()
	if err != nil {
		t.Errorf("Temp directory not created %v", err)
	}

	err = CleanupTempDir((dir))
	if err != nil {
		t.Errorf("Unable to cleanup temp directory, %v", err)
	}
}

func TestCreateTempDir_Negative(t *testing.T) {
	/* TODO: Add negative test. */
}

func TestCleanupTempDir_Positive(t *testing.T) {
	dir, err := CreateTempDir()
	if err != nil {
		t.Errorf("Temp directory not created %v", err)
	}

	err = CleanupTempDir((dir))
	if err != nil {
		t.Errorf("Unable to cleanup temp directory, %v", err)
	}

	_, err = os.Stat(dir)
	if err == nil {
		t.Errorf("Temp directory still exists")
	} else if !os.IsNotExist(err) {
		t.Errorf("Error checking if directory exists: %v", err)
	}
}

func TestCleanupTempDir_Negative(t *testing.T) {
	/* TODO: Add negative test. */
}
