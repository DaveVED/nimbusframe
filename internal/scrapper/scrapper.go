package scrapper

import (
	"encoding/csv"
	"io/ioutil"
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

	resources, err := extractResources(files)
	if err != nil {
		return nil, err
	}

	err = writeNames(resources)
	if err != nil {
		return nil, err
	}

	return []string{}, nil
}

func writeNames(resources map[string]string) error {
	/* TODO: If file exits check, handle that case */
	csvFile, err := os.Create("names/resource_names.csv")
	if err != nil {
		return err
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	writer.Write([]string{"SDKResource", "Name"})

	for key, value := range resources {
		writer.Write([]string{key, value})
	}

	return nil
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
				tempFile := filepath.Join(dir, fileName)
				files = append(files, tempFile)
			}
		}
	}

	return files, nil
}

func extractResources(files []string) (map[string]string, error) {
	resources := make(map[string]string)

	for _, entry := range files {
		fileContent, err := ioutil.ReadFile(entry)
		if err != nil {
			return nil, err
		}

		lines := strings.Split(string(fileContent), "\n")
		for _, line := range lines {
			if strings.Contains(line, "// @SDKResource") {
				start := strings.Index(line, "(")
				end := strings.Index(line, ")")
				tempLine := line[start+1 : end]
				values := strings.Split(tempLine, ",")
				rawResource := values[0]
				resource := rawResource[1 : len(rawResource)-1]

				name := ""
				if len(values) == 2 {
					fullValue := values[1]
					name = fullValue[7 : len(fullValue)-1]
				}

				resources[resource] = name
			}
		}
	}

	return resources, nil
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
