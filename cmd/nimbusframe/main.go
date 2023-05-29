package main

import (
	"github.com/DaveVED/nimbusframe/internal/utils"
)

func main() {
	tempDir, err := utils.CloneGithubRepo("https://github.com/hashicorp/terraform-provider-aws.git")
	if err != nil {
		return
	}

	utils.CleanupTempDir(tempDir)
}
