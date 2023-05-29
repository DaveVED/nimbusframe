package scrapper

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func FindResources(dir string) []string {
	servicesPath := dir + "/internal/service/ec2"
	files, err := searchDir(servicesPath)
	fmt.Println(files)
	fmt.Println(err)
	return []string{}
}

func searchDir(dir string) ([]string, error) {
	/* TODO: If not a *.go file, check if it's a dir. If it's a dir serach it... Add recursive logic. */
	files := []string{}
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		fileName := e.Name()
		if isGoFile(fileName) {
			files = append(files, fileName)
		}
	}
	return files, nil
}

func extractResource(path string) string {
	/* TODO: Write code to extract SDKResource from Terraform go file. */
	return ""
}

func isGoFile(file string) bool {
	components := strings.Split(file, ".")

	if len(components) == 0 {
		return false
	}

	extension := components[len(components)-1]
	return extension == "go"
}
