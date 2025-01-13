// Package paradirectories provides custom types for structuring Para Method's Directories.
package paradirectories

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
		// Name of the Directory.
		Name string `toml:"name"`
		// Content to be written to the README.md file inside every para folder.
		ReadMeContent string `toml:"readme_content"`
	}

	// ParaStructure contains all the necessary information for the script to work.
	ParaStructure struct {
		// An array of ParaDirectories.
		Directories []ParaDirectory `toml:"directories"`
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

	return nil
}

// GenerateParaDirectory creates the directory for the specified paraDirectory.
func GenerateParaDirectory(dir ParaDirectory, baseDir string) error {
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
func ShowFileTree(baseDir string, paraDirectories []ParaDirectory) string {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Fatal(err)
	}

	buf := bytes.Buffer{} // We are writing everything on here.

	fmt.Fprintln(&buf, "")
	fmt.Fprintln(&buf, strings.Repeat("=", width))
	fmt.Fprintln(&buf, "")
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

	fmt.Fprintln(&buf, "")
	fmt.Fprintln(&buf, strings.Repeat("=", width))
	fmt.Fprintln(&buf, "")

	return buf.String() // Returns everything that was written on the buffer 
}
