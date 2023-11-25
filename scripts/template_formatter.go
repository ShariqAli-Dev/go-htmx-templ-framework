package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	baseTemplatePath, err := filepath.Abs("views/partials/base_html_templ.go")
	if err != nil {
		log.Fatal(err)
	}
	content, err := os.ReadFile(baseTemplatePath)
	if err != nil {
		log.Fatal(err)
	}
	fileContent := string(content)

	searchString := "main.js"
	replacementString := getBuiltJS("index-")
	fileContent = strings.Replace(fileContent, searchString, replacementString, 1)
	err = os.WriteFile(baseTemplatePath, []byte(fileContent), 0644)
	if err != nil {
		log.Fatal(err)
	}

}

func getBuiltJS(fileBase string) string {
	assetsDirectoryPath, err := filepath.Abs("web/dist/assets")
	if err != nil {
		log.Fatal(err)
	}
	dir, err := os.Open(assetsDirectoryPath)
	if err != nil {
		log.Fatal(err)
	}
	defer dir.Close()

	files, err := dir.Readdirnames(-1)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if strings.Contains(file, fileBase) {
			return file
		}
	}
	log.Fatalf("could not locate a file with the base %s", fileBase)
	return ""
}
