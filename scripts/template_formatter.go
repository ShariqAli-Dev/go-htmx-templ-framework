package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	routesPath, err := filepath.Abs("views")
	if err != nil {
		log.Fatal(err)
	}
	if err := filepath.Walk(routesPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && path != routesPath {
			return filepath.SkipDir
		}

		if strings.HasSuffix(info.Name(), ".go") {
			fileName := info.Name()
			fileNameWithoutSuffix := fileName[:len(fileName)-9]
			searchFileName := getBuiltJS(fileNameWithoutSuffix + "-")
			if searchFileName == "" {
				return nil
			}

			template, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			templateContent := string(template)

			searchString := fmt.Sprintf(`HTML("%s")`, fileNameWithoutSuffix)
			replacementString := fmt.Sprintf(`HTML("%s")`, searchFileName)

			templateContent = strings.Replace(templateContent, searchString, replacementString, 1)
			if err := os.WriteFile(path, []byte(templateContent), 0644); err != nil {
				log.Fatal(err)
			}

			return nil
		}
		return nil
	}); err != nil {
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
	return ""
}
