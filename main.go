// Package main provides a utility for generating a PARA directory structure.
// The PARA method (Projects, Areas, Resources, and Archive) is a framework for
// organizing files and notes effectively. This package allows you to generate
// the structure in a specified base directory and optionally preview it
// beforehand.
package main

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	p "github.com/Kacaii/para-structure-generator/paraDirectories"

	"github.com/BurntSushi/toml"
)

// Constants for stylizing the text output.
const (
	greenColor string = "\x1b[32m"
	resetColor string = "\x1b[0m"
)

var (

	// config is embedded into the binary, so its always available.
	//
	//go:embed config.toml
	config string

	// Define the "create" subcommand and its flags
	createCmd = flag.NewFlagSet("create", flag.ExitOnError)

	// baseDir represents the directory where the structure will be generated
	baseDir = createCmd.String("b", ".", "Base directory for generating the structure")

	// Define the global -h flag
	printHelp = flag.Bool("h", false, "Prints the help message")
)

// main is the entry point of the program.
func main() {
	flag.Parse() // Parse the global flags

	var paraStructure p.ParaStructure
	if err := toml.Unmarshal([]byte(config), &paraStructure); err != nil {
		log.Fatal("Error parsing TOML file:", err)
	}

	if *printHelp {
		flag.Usage()
		createCmd.Usage()
		os.Exit(0)
	}

	if len(os.Args) < 2 {
		log.Fatal("Expected 'create' subcommand")
	}

	// Checks for subcommand
	switch os.Args[1] {
	case "create":
		createCmd.Parse(os.Args[2:]) // parses the flags for the "create" subcommand
		handleCreate(*baseDir, paraStructure)
	default:
		fmt.Println("Unknown subcommand")
		os.Exit(1)
	}
}

// handleCreate validades the base directory and generates the file structure 󰔱
func handleCreate(baseDir string, paraStructure p.ParaStructure) {
	if err := ValidateBaseDir(baseDir); err != nil {
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
func WriteReadme(dir p.ParaDirectory, baseDirectory string) error {
	// First we need to path to the README file.
	filePath := filepath.Join(baseDirectory, dir.Name, "README.md")

	// Then we write the contents to it.
	if err := os.WriteFile(filePath, []byte(dir.ReadMeContent), os.ModePerm); err != nil {
		return err
	}

	return nil
}

// GenerateParaDirectory creates the directory for the specified paraDirectory.
func GenerateParaDirectory(dir p.ParaDirectory, baseDir string) error {
	pathToDirectory := filepath.Join(baseDir, dir.Name)
	return os.MkdirAll(pathToDirectory, os.ModePerm)
}

// ValidateBaseDir checks if the provided base directory is valid and accessible.
func ValidateBaseDir(baseDir string) error {
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
func ShowFileTree(baseDir string, paraDirectories []p.ParaDirectory) string {
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
