package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type paraFolder struct {
	folderNumber string
	folderName   string
	fileName     string
	fileContent  string
}

func main() {
	baseDir := "PARA"

	paraStructure := []paraFolder{
		{
			"01",
			"PROJECTS",
			"PROJECTS.md",
			`## 01 - PROJECTS

Contains files and resources related to active projects. Each project is tracked and managed here.`,
		},
		{
			"02",
			"AREAS",
			"AREAS.md",
			`## 02 - AREAS

Holds information about ongoing responsibilities and areas of focus.`,
		},
		{
			"03",
			"RESOURCES",
			"RESOURCES.md",
			`## 03 - RESOURCES

A collection of materials, references, and other helpful resources.`,
		},
		{
			"04",
			"ARQUIVE",
			"ARQUIVE.md",
			`## 04 - ARQUIVE

Archive for completed projects, inactive areas, or other no-longer-active items.`,
		},
	}

	for _, folder := range paraStructure {
		// Creates the directories
		err := os.MkdirAll(filepath.Join(baseDir, folder.folderName), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}

		// Create and write to file
		filePath := filepath.Join(baseDir, folder.folderName, folder.fileName)

		err = os.WriteFile(filePath, []byte(folder.fileContent), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s Created successfully!\n", filePath)
	}

	fmt.Println("PARA structure generated successfully using Golang! 󱜙")
}
