package utils

import (
	"os"
	"testing"
)

func TestCloneGithubRepo_Positive(t *testing.T) {
	dir, err := CloneGithubRepo("https://github.com/hashicorp/terraform-provider-aws.git")
	if err != nil {
		t.Errorf("Unable to clone the repository: %v", err)
	}

	// Check if the directory exists
	_, err = os.Stat(dir)
	if os.IsNotExist(err) {
		t.Errorf("Clone directory does not exist: %v", err)
	} else if err != nil {
		t.Errorf("Error checking if clone directory exists: %v", err)
	}

	// Check if the directory is not empty
	files, err := os.ReadDir(dir)
	if err != nil {
		t.Errorf("Error reading clone directory: %v", err)
	} else if len(files) == 0 {
		t.Errorf("Clone directory is empty")
	}

	err = CleanupTempDir((dir))
	if err != nil {
		t.Errorf("Unable to cleanup temp directory, %v", err)
	}
}

func TestCloneGithubRepo_Negative(t *testing.T) {
	/* TODO: Add negative test. */
}
