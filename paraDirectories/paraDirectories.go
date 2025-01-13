// Package paradirectories provides custom types for structuring Para Method's Directories.
package paradirectories

import (
	"os"
	"path/filepath"
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
