package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const GREEN string = "\x1b[32m"

const (
	projectsDesc  string = "Stores notes and files for active, time-bound tasks or deliverables."
	areasDesc     string = "Contains ongoing responsibilities or areas of interest."
	resourcesDesc string = "Holds general reference materials and reusable templates."
	arquiveDesc          = "Keeps inactive projects and outdated resources for future reference."
)

type paraDirectory struct {
	name          string
	readMeContent string
}

func main() {
	var baseDir string // Allows the user to select a base directory when generating the PARA structure
	flag.StringVar(&baseDir, "dir", ".", "Select a path  to the base directory for the file structure to be generated 󰞋")
	flag.Parse() // Allows the flags to be accessed by the program

	err := validateBaseDir(baseDir)
	if err != nil {
		log.Fatalf("Invalid base directory: %v", err)
	}

	paraStructure := []paraDirectory{
		{
			"01 PROJECTS",
			projectsDesc,
		},
		{
			"02 AREAS",
			areasDesc,
		},
		{
			"03 RESOURCES",
			resourcesDesc,
		},
		{
			"04 ARQUIVE",
			arquiveDesc,
		},
	}

	fmt.Printf("Generating PARA structure in: %s \n", baseDir) // Feedback message for the user.

	err = generateParaDirectories(paraStructure, baseDir)
	if err != nil {
		log.Fatal("Failed to generate directories", err)
	}

	err = writeFileContent(paraStructure, baseDir)
	if err != nil {
		log.Fatal("Failed to write content to README files", err)
	}

	fmt.Println(GREEN + "PARA Structure Generated Successfully Using Golang! 󱜙  ") // All done! 
}

// Writes content to the PARA Files
func writeFileContent(paraStructure []paraDirectory, baseDirectory string) error {
	for _, dir := range paraStructure {
		// Generating a path for every README file
		filePath := filepath.Join(baseDirectory, dir.name, "README.md")

		// Writing contents to the README files
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
		// Creates a path for every directory of the structure
		pathToDirectory := filepath.Join(baseDir, dir.name)

		// Generates a directory for every part of the structure
		err := os.MkdirAll(pathToDirectory, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

// Returns an error if the path provided dont exist or if is not a directory.
func validateBaseDir(baseDir string) error {
	// Information about the path provided
	info, err := os.Stat(baseDir)
	if os.IsNotExist(err) {
		return fmt.Errorf("path does not exist: %s", baseDir)
	}

	if err != nil {
		return fmt.Errorf("unable to access path: %v", err)
	}

	if !info.IsDir() {
		return fmt.Errorf("path is not a directory: %s", baseDir)
	}

	return nil
}
