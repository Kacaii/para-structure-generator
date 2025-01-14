// Package main provides a utility for generating a PARA directory structure.
// The PARA method (Projects, Areas, Resources, and Archive) is a framework for
// organizing files and notes effectively. This package allows you to generate
// the structure in a specified base directory.
package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	para "github.com/Kacaii/para-structure-generator/paraDirectories"

	"github.com/BurntSushi/toml"
)

// configFile is a TOML file embedded into the binary, so its always available. 
//
//go:embed config.toml
var configFile string

// Constants for stylizing the text output.
const (
	greenColor string = "\x1b[32m"
	resetColor string = "\x1b[0m"
)

var (

	// Define the "create" subcommand and its flags
	createCmd = flag.NewFlagSet("create", flag.ExitOnError)

	// baseDir represents the directory where the structure will be generated
	baseDir = createCmd.String("b", ".", "Base directory for generating the structure")

	// Define the global -h flag
	printHelp = flag.Bool("h", false, "Prints the help message")
)

// main is the entry point of the program.
func main() {
	flag.Parse() // Parse the global flags 

	var paraStructure para.ParaStructure
	// Parsing the embedded config file.
	if err := toml.Unmarshal([]byte(configFile), &paraStructure); err != nil {
		log.Fatal("Error parsing TOML file:", err)
	}

	// Handle printing the help message 󰞋
	if *printHelp {
		flag.Usage()      // Global Flags 
		createCmd.Usage() // "Create" subcommand flags 
		os.Exit(0)        // Exits the program
	}

	// User needs to provide a subcommand
	if len(os.Args) < 2 {
		log.Fatal("Expected 'create' subcommand")
	}

	// Checks for subcommand
	switch os.Args[1] {
	case "create":
		createCmd.Parse(os.Args[2:]) // Parse the flags for the "create" subcommand 
		handleCreate(*baseDir, paraStructure)
	default:
		fmt.Println("Unknown subcommand") // In case the user doesnt pass a subcommand.
		os.Exit(1)
	}
}

// handleCreate validades the base directory and generates the file structure 󰔱
func handleCreate(baseDir string, paraStructure para.ParaStructure) {
	if err := para.ValidateBaseDir(baseDir); err != nil {
		log.Fatalln("Invalid base directory:", err)
	}

	if baseDir == "." {
		fmt.Println("Generating PARA structure in the current directory 󰣞")
	} else {
		fmt.Println("Generating PARA structure in:", baseDir, "󰣞")
	}

	var wg sync.WaitGroup
	for _, dir := range paraStructure.Directories {
		// Add one (1) to the waitGroup for every directory in the structure.
		wg.Add(1)

		// Spawning goroutines
		go func() {
			defer wg.Done()

			// First we generate the directories.
			if err := para.GenerateParaDirectory(dir, baseDir); err != nil {
				log.Println("Failed to create directory:", dir.Name, err.Error())
			}

			// Then we create and write the README files.
			if err := para.WriteReadme(dir, baseDir); err != nil {
				log.Println("Failed to write README for:", dir.Name, err.Error())
			}
		}()
	}

	wg.Wait() // Waiting for all goroutines to finish.

	// And showing the result at te end.
	fmt.Println(para.ShowFileTree(baseDir, paraStructure.Directories))
	fmt.Println(greenColor + "PARA Structure Generated Successfully Using Golang! 󱜙  " + resetColor) // All done! 
}
