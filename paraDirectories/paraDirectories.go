// Package paradirectories provides custom types for structuring Para Method's Directories.
package paradirectories

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
