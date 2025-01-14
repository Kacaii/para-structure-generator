/*
Generate-para provides a utility for generating a PARA method structure.

The PARA method (Projects, Areas, Resources, and Archive) is a framework for
organizing files and notes effectively. This package allows you to generate
the structure in a specified base directory. It includes default README files
summarizing whats each directory is used for.

Usage:

	generate-para [flags] [subcommand] [subflags] [path ...]

The flags are:

	-h

	    Prints the help message to the console.

The subcommands are:

	create [flags] [path ...]

	    Generates the PARA structure in the specified path.
	    Defaults to the current directory.

The subflags for create are:

	-b

	    Specify the base directory for the structure to be generated.
	    Defaults to the current directory.
*/
package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	p "github.com/Kacaii/para-structure-generator/paramethod"

	"github.com/BurntSushi/toml"
)

// configFile is a TOML file embedded into the binary, so its always available. 
//
//		──────────────────
//		  ├─  bin
//	    │
//		  ├─  config.toml
//		  ├─  main.go
//	    │
//		  ├─  go.mod
//		  └─  go.sum
//		──────────────────
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
	//
	// Usage:
	//
	//  generate-para create [flags] [path...]
	createCmd = flag.NewFlagSet("create", flag.ExitOnError)

	// baseDir represents the directory where the structure will be generated
	//
	// Usage:
	//
	//  generate-para create -b [base-directory]
	baseDir = createCmd.String("b", ".", "Base directory for generating the structure")

	// Define the global -h flag
	//
	// Usage:
	//
	//  generate-para -h
	printHelp = flag.Bool("h", false, "Prints the help message")
)

// main is the entry point of the program.
func main() {
	flag.Parse() // Parse the global flags 

	paraMethod := parseTOML()

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
		handleCreate(*baseDir, paraMethod)
	default:
		fmt.Println("Unknown subcommand") // In case the user doesnt pass a subcommand.
		os.Exit(1)
	}
}

func parseTOML() p.ParaMethod {
	var config p.ParaMethod
	// Parsing the embedded config file.
	if err := toml.Unmarshal([]byte(configFile), &config); err != nil {
		log.Fatal("Error parsing TOML file:", err)
	}

	return config
}

// handleCreate validades the base directory and generates the file structure 󰔱
func handleCreate(baseDirectory string, paraStructure p.ParaMethod) {
	if err := p.ValidateBaseDir(baseDirectory); err != nil {
		log.Fatalln("Invalid base directory:", err)
	}

	if baseDirectory == "." {
		fmt.Println("Generating PARA structure in the current directory 󰣞")
	} else {
		fmt.Println("Generating PARA structure in:", baseDirectory, "󰣞")
	}

	var wg sync.WaitGroup
	for _, paraDir := range paraStructure.Directories {
		// Add one (1) to the waitGroup for every directory in the structure.
		wg.Add(1)

		// Spawning goroutines
		go func() {
			defer wg.Done()

			// First we generate the directories.
			if err := p.GenerateParaDirectory(paraDir, baseDirectory); err != nil {
				log.Println("Failed to create directory:", paraDir.Name, err.Error())
			}

			// Then we create and write the README files.
			if err := p.WriteReadme(paraDir, baseDirectory); err != nil {
				log.Println("Failed to write README for:", paraDir.Name, err.Error())
			}
		}()
	}

	wg.Wait() // Waiting for all goroutines to finish.

	// And showing the result at te end.
	fmt.Println(p.ShowFileTree(baseDirectory, paraStructure.Directories))
	fmt.Println(greenColor + "PARA Structure Generated Successfully Using Golang! 󱜙  " + resetColor) // All done! 
}
