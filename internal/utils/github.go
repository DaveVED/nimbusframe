package utils

import (
	"os/exec"
)

func CloneGithubRepo(source string) (string, error) {
	/* We first have to create a temp directory. This will avoid, repository already exists */
	tempDir, err := CreateTempDir()
	if err != nil {
		return "", err
	}

	/* Clone down the repo, we just use exec command for now */
	cmd := exec.Command("git", "clone", source, tempDir)
	err = cmd.Run()
	if err != nil {
		return "", err
	}

	/* Give them clone destination */
	return tempDir, nil
}
