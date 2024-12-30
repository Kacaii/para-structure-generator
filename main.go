package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
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
		{"01 PROJECTS", projectsDesc},
		{"02 AREAS", areasDesc},
		{"03 RESOURCES", resourcesDesc},
		{"04 ARQUIVE", arquiveDesc},
	}

	if baseDir == "." {
		fmt.Println("Generating PARA structure in the current directory 󰣞")
	} else {
		fmt.Println("Generating PARA structure in:", baseDir, "󰣞")
	}

	var wg sync.WaitGroup
	for _, dir := range paraStructure {
		// Add one (1) to the waitGroups for every directory in the structure
		wg.Add(1)

		// Spawn goroutines  to generate the whole file structure
		go func(d paraDirectory) {
			defer wg.Done()

			err = generateParaDirectory(d, baseDir)
			if err != nil {
				log.Println("Failed to create directory:", d.name, err.Error())
			}

			err = writeReadme(d, baseDir)
			if err != nil {
				log.Println("Failed to write README for:", d.name, err.Error())
			}
		}(dir) // Passing every directory from the structure as an arguemnt to the function
	}

	wg.Wait() // Waiting for every file to be written

	fmt.Println(GREEN + "PARA Structure Generated Successfully Using Golang! 󱜙  ") // All done! 
}

// Writes content to the PARA Files
func writeReadme(dir paraDirectory, baseDirectory string) error {
	filePath := filepath.Join(baseDirectory, dir.name, "README.md")
	// Writing contents to the README files
	err := os.WriteFile(filePath, []byte(dir.readMeContent), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// Generates the necessary Directories for the structure: PROJECTS, AREAS, RESOURCES and ARQUIVE
func generateParaDirectory(dir paraDirectory, baseDir string) error {
	pathToDirectory := filepath.Join(baseDir, dir.name)
	return os.MkdirAll(pathToDirectory, os.ModePerm)
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
