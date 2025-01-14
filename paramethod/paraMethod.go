// Package paramethod provides custom types and functions for structuring PARA Method's Directories.
package paramethod

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/term"
)

type (
	// ParaDirectory defines a directory in the PARA structure with a name and description.
	ParaDirectory struct {
		Name          string `toml:"name"`           // Name of the Directory.
		ReadMeContent string `toml:"readme_content"` // Content to be written to the README.md file inside every PARA directory.
	}

	// ParaMethod contains all the necessary information for the script to work.
	ParaMethod struct {
		Directories []ParaDirectory `toml:"directories"` // [ PROJECTS, AREAS, RESOURCES, ARQUIVE ]
	}
)

// WriteReadme writes the content to a README.md file in the specified directory.
func WriteReadme(dir ParaDirectory, baseDirectory string) error {
	// First we need to path to the README file.
	filePath := filepath.Join(baseDirectory, dir.Name, "README.md")

	// Then we write the contents to it.
	if err := os.WriteFile(filePath, []byte(dir.ReadMeContent), os.ModePerm); err != nil {
		return err
	}

	return nil // If nothing goes wrong, return nil.
}

// GenerateParaDirectory creates the directory for the specified paraDirectory.
func GenerateParaDirectory(dir ParaDirectory, baseDir string) error {
	// Path to the directory that we are generating.
	pathToDirectory := filepath.Join(baseDir, dir.Name)

	// Returns nil if no error ocurred.
	return os.MkdirAll(pathToDirectory, os.ModePerm)
}

// ValidateBaseDir checks if the provided base directory is valid and accessible.
func ValidateBaseDir(baseDir string) error {
	// Information about the path provided.
	info, err := os.Stat(baseDir)

	// If it doesnt exists.
	if os.IsNotExist(err) {
		return fmt.Errorf("selected path does not exist: %s", baseDir)
	}

	// If path isnt accessible.
	if err != nil {
		return fmt.Errorf("unable to access path: %v", err)
	}

	// If path is a file.
	if !info.IsDir() {
		return fmt.Errorf("path is not a directory: %s", baseDir)
	}

	return nil // If nothing goes wrong, return nil.
}

// ShowFileTree returns a string representation of the PARA file structure.
func ShowFileTree(baseDir string, paraDirectories []ParaDirectory) string {
	// terminalWidth stores the width of the terminal.
	terminalWidth, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Fatal(err)
	}

	// Buffer to store the bytes that are going to be displayed in the terminal. 󰜦
	buf := bytes.Buffer{} //

	fmt.Fprintln(&buf, "")
	fmt.Fprintln(&buf, strings.Repeat("=", terminalWidth))
	fmt.Fprintln(&buf, "")
	fmt.Fprintln(&buf, baseDir+"/") // Base directory.
	fmt.Fprintln(&buf, "│")

	// Previews the file tree showing each of its directories
	for i, dir := range paraDirectories {
		if i+1 != len(paraDirectories) {
			fmt.Fprintln(&buf, "├──", dir.Name+"/")    // Directories 1, 2 and 3.
			fmt.Fprintln(&buf, "│   └──", "README.md") // Every directory has a README file.
			fmt.Fprintln(&buf, "│")
		} else {
			fmt.Fprintln(&buf, "└──", dir.Name+"/")    // Final directory.
			fmt.Fprintln(&buf, "    └──", "README.md") // README for the final directory.
		}
	}

	fmt.Fprintln(&buf, "")
	fmt.Fprintln(&buf, strings.Repeat("=", terminalWidth))

	return buf.String() // Returns everything that was written on the buffer. 
}
