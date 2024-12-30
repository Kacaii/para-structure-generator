package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type paraFolder struct {
	folderName  string
	fileContent string
}

// TODO: Add option to specify base directory
var baseDir string = "PARA"

func main() {
	paraStructure := []paraFolder{
		{
			"01 PROJECTS",
			"Stores notes and files for active, time-bound tasks or deliverables.",
		},
		{
			"02 AREAS",
			"Contains ongoing responsibilities or areas of interest.",
		},
		{
			"03 RESOURCES",
			"Holds general reference materials and reusable templates.",
		},
		{
			"04 ARQUIVE",
			"Keeps inactive projects and outdated resources for future reference.",
		},
	}

	for _, folder := range paraStructure {
		// Creates the directories
		err := os.MkdirAll(filepath.Join(baseDir, folder.folderName), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}

		// Create and write to file
		filePath := filepath.Join(baseDir, folder.folderName, "README.md")

		err = os.WriteFile(filePath, []byte(folder.fileContent), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("PARA Structure Generated Successfully Using Golang! 󱜙 ")
}
