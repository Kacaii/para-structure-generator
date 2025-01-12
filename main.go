// Package main provides a utility for generating a PARA directory structure.
// The PARA method (Projects, Areas, Resources, and Archive) is a framework for
// organizing files and notes effectively. This package allows you to generate
// the structure in a specified base directory and optionally preview it
// beforehand.
package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/BurntSushi/toml"
)

// Its embedded into the binary, so its always available.
//
//go:embed config.toml
var config string

// Constants for stylizing the text output.
const (
	greenColor string = "\x1b[32m"
	resetColor string = "\x1b[0m"
)

// ParaDirectory defines a directory in the PARA structure with a name and description.
type ParaDirectory struct {
	Name          string `toml:"name"`           // Name of the Directory
	ReadMeContent string `toml:"readme_content"` // Content for the README.md file
}

// ParaStructure contains all the necessary information for the script to work
type ParaStructure struct {
	Directories []ParaDirectory `toml:"directories"` // Needs to be an Exported field so the other libraries can detect it
}

// main is the entry point of the program.
func main() {
	var paraStructure ParaStructure

	if err := toml.Unmarshal([]byte(config), &paraStructure); err != nil {
		log.Fatal("Error parsing TOML file:", err)
	}

	if len(os.Args) != 2 {
		log.Fatal("Please select a base directory")
	}

	baseDir := os.Args[1]

	if err := validateBaseDir(baseDir); err != nil {
		log.Fatalln("Invalid base directory:", err)
	}

	if baseDir == "." {
		fmt.Println("Generating PARA structure in the current directory 󰣞")
	} else {
		fmt.Println("Generating PARA structure in:", baseDir, "󰣞")
	}

	var wg sync.WaitGroup
	for _, dir := range paraStructure.Directories {
		// Add one (1) to the waitGroups for every directory in the structure.
		wg.Add(1)

		// Spawning goroutines
		go func() {
			defer wg.Done()

			if err := GenerateParaDirectory(dir, baseDir); err != nil {
				log.Println("Failed to create directory:", dir.Name, err.Error())
			}

			if err := WriteReadme(dir, baseDir); err != nil {
				log.Println("Failed to write README for:", dir.Name, err.Error())
			}
		}()
	}

	wg.Wait() // Waiting for all goroutines to finish.

	fmt.Println("")
	fmt.Println(ShowFileTree(baseDir, paraStructure.Directories))
	fmt.Println("")
	fmt.Println(greenColor + "PARA Structure Generated Successfully Using Golang! 󱜙  " + resetColor) // All done! 
}

// WriteReadme writes the content to a README.md file in the specified directory.
func WriteReadme(dir ParaDirectory, baseDirectory string) error {
	// First we need to path to the README file.
	filePath := filepath.Join(baseDirectory, dir.Name, "README.md")

	// Then we write the contents to it.
	if err := os.WriteFile(filePath, []byte(dir.ReadMeContent), os.ModePerm); err != nil {
		return err
	}

	return nil
}

// GenerateParaDirectory creates the directory for the specified paraDirectory.
func GenerateParaDirectory(dir ParaDirectory, baseDir string) error {
	pathToDirectory := filepath.Join(baseDir, dir.Name)
	return os.MkdirAll(pathToDirectory, os.ModePerm)
}

// validateBaseDir checks if the provided base directory is valid and accessible.
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

	return nil // If nothing goes wrong, return nil.
}

// ShowFileTree returns a string representation of the PARA file structure.
func ShowFileTree(baseDir string, paraDirectories []ParaDirectory) string {
	buf := bytes.Buffer{} // We are writing everything on here.

	fmt.Fprintln(&buf, baseDir+"/") // Base directory.
	fmt.Fprintln(&buf, "│")

	// Previews the file tree showing each of its directories
	for i, dir := range paraDirectories {
		if i+1 != len(paraDirectories) {
			fmt.Fprintln(&buf, "├──", dir.Name+"/")
			fmt.Fprintln(&buf, "│   └──", "README.md") // Every directory has a README file
			fmt.Fprintln(&buf, "│")
		} else {
			fmt.Fprintln(&buf, "└──", dir.Name+"/")    // Final directory
			fmt.Fprintln(&buf, "    └──", "README.md") // README for the final directory
		}
	}

	return buf.String() // Returns everything that was written on the buffer 
}
