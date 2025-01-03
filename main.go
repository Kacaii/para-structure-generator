package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
)

const (
	GREEN string = "\x1b[32m"
	RESET string = "\x1b[0m"
)

const (
	projectsDesc  string = "Stores notes and files for active, time-bound tasks or deliverables."
	areasDesc     string = "Contains ongoing responsibilities or areas of interest."
	resourcesDesc string = "Holds general reference materials and reusable templates."
	arquiveDesc   string = "Keeps inactive projects and outdated resources for future reference."
)

type paraDirectory struct {
	name          string
	readMeContent string
}

func main() {
	paraStructure := []paraDirectory{
		{"01 PROJECTS", projectsDesc},
		{"02 AREAS", areasDesc},
		{"03 RESOURCES", resourcesDesc},
		{"04 ARQUIVE", arquiveDesc},
	}

	baseDir, previewTree := handleFlags() // Allows the flags to be accessed by the program

	if previewTree {
		fmt.Println(showFileTree(baseDir, paraStructure)) // Shows the file tree for testing purpose.
		os.Exit(0)
	}

	if err := validateBaseDir(baseDir); err != nil {
		log.Fatalln("Invalid base directory:", err)
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

		// Spawning goroutines
		go func() {
			defer wg.Done()

			if err := generateParaDirectory(dir, baseDir); err != nil {
				log.Println("Failed to create directory:", dir.name, err.Error())
			}

			if err := writeReadme(dir, baseDir); err != nil {
				log.Println("Failed to write README for:", dir.name, err.Error())
			}
		}() // Passing every directory from the structure as an arguemnt to the function
	}

	wg.Wait() // Waiting for every file to be written

	fmt.Println("")
	fmt.Println(showFileTree(baseDir, paraStructure))
	fmt.Println("")
	fmt.Println(GREEN + "PARA Structure Generated Successfully Using Golang! 󱜙  " + RESET) // All done! 
}

func handleFlags() (baseDir string, previewTree bool) {
	flag.StringVar(&baseDir, "dir", ".", "Base directory for generating the File Structure")
	flag.BoolVar(&previewTree, "tree", false, "Preview the File Structure without creating it")

	flag.Parse()

	return baseDir, previewTree
}

// Writes content to the PARA Files
func writeReadme(dir paraDirectory, baseDirectory string) error {
	// First we need to path to the README file
	filePath := filepath.Join(baseDirectory, dir.name, "README.md")

	// Then we write the contents to it
	if err := os.WriteFile(filePath, []byte(dir.readMeContent), os.ModePerm); err != nil {
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
		return fmt.Errorf("selected path does not exist: %s", baseDir)
	}

	if err != nil {
		return fmt.Errorf("unable to access path: %v", err)
	}

	if !info.IsDir() {
		return fmt.Errorf("path is not a directory: %s", baseDir)
	}

	return nil
}

func showFileTree(baseDir string, paraStructure []paraDirectory) string {
	buf := bytes.Buffer{} // We are writing everything on here

	fmt.Fprintln(&buf, baseDir+"/") // Writtes the base directory
	fmt.Fprintln(&buf, "│")

	// Previews the file tree showing each of its directories
	for i, dir := range paraStructure {
		if i+1 != len(paraStructure) {
			fmt.Fprintln(&buf, "├──", dir.name+"/")
			fmt.Fprintln(&buf, "│   └──", "README.md") // Every directory has a README file
			fmt.Fprintln(&buf, "│")
		} else {
			fmt.Fprintln(&buf, "└──", dir.name+"/")    // Final directory
			fmt.Fprintln(&buf, "    └──", "README.md") // README for the final directory
		}
	}

	return buf.String() // Returns everything that was written on the buffer 
}
