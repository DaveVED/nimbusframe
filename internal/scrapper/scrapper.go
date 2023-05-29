package scrapper

import (
	"fmt"
	"os"
	"strings"
)

func FindResources(dir string) ([]string, error) {
	servicesPath := dir + "/internal/service/ec2"
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
	/* TODO: If not a *.go file, check if it's a dir. If it's a dir serach it... Add recursive logic. */
	files := []string{}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, e := range entries {
		fileName := e.Name()
		if isGoFile(fileName) {
			files = append(files, fileName)
		}
	}
	return files, nil
}

func extractResources(files []string) ([]string, error) {
	/* TODO: Write code to extract SDKResourcse from Terraform go files. */
	return []string{}, nil
}

func isGoFile(file string) bool {
	components := strings.Split(file, ".")

	if len(components) == 0 {
		return false
	}

	extension := components[len(components)-1]
	return extension == "go"
}
