package scrapper

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func FindResources(dir string) ([]string, error) {
	servicesPath := dir + "/internal/service"
	files, err := searchDir(servicesPath)
	if err != nil {
		return nil, err
	}
	fmt.Println(files)

	resources, err := extractResources(files)
	if err != nil {
		return nil, err
	}
	fmt.Println(resources)

	return resources, nil
}

func searchDir(dir string) ([]string, error) {
	files := []string{}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			subDir := filepath.Join(dir, entry.Name())
			subFiles, err := searchDir(subDir)
			if err != nil {
				return nil, err
			}
			files = append(files, subFiles...)
		} else {
			fileName := entry.Name()
			/* TODO: And not a _test file -- change isGoFile to isValidGoFile */
			if isGoFile(fileName) {
				files = append(files, fileName)
			}
		}
	}

	return files, nil
}

func extractResources(files []string) ([]string, error) {
	/* TODO: Write code to extract SDKResourcse from Terraform go files. */
	return []string{}, nil
}

func isGoFile(file string) bool {
	/* TODO: And not a _test file -- change isGoFile to isValidGoFile */
	/* TODO: Find common files, and also do not add those */
	components := strings.Split(file, ".")

	if len(components) == 0 {
		return false
	}

	extension := components[len(components)-1]
	return extension == "go"
}
