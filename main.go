package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const GREEN string = "\x1b[32m"

type paraDirectory struct {
	name          string
	readMeContent string
}

func main() {
	var baseDir string // Allows the user to select a base directory to generate the structure
	flag.StringVar(&baseDir, "dir", ".", "Select a base directory for the structure to be generated")
	flag.Parse() // Allows the flags to be accessed by the program

	paraStructure := []paraDirectory{
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

	fmt.Printf("Generating PARA structure in: %s \n", baseDir)

	err := generateParaDirectories(paraStructure, baseDir)
	if err != nil {
		log.Fatal("Failed to generate directories", err)
	}

	err = writeFileContent(paraStructure, baseDir)
	if err != nil {
		log.Fatal("Failed to write content to README files", err)
	}

	fmt.Println(GREEN + "PARA Structure Generated Successfully Using Golang! 󱜙 ") // All done!
}

// Writes content to the PARA Files
func writeFileContent(paraStructure []paraDirectory, baseDirectory string) error {
	for _, dir := range paraStructure {
		filePath := filepath.Join(baseDirectory, dir.name, "README.md")

		err := os.WriteFile(filePath, []byte(dir.readMeContent), os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

// Generates the necessary Directories for the structure: PROJECTS, AREAS, RESOURCES and ARQUIVE
func generateParaDirectories(structure []paraDirectory, baseDir string) error {
	for _, dir := range structure {
		pathToDirectory := filepath.Join(baseDir, dir.name)

		err := os.MkdirAll(pathToDirectory, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}
