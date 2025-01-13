package paradirectories

// ParaDirectory defines a directory in the PARA structure with a name and description.
type ParaDirectory struct {
	Name          string `toml:"name"`           // Name of the Directory
	ReadMeContent string `toml:"readme_content"` // Content for the README.md file
}

// ParaStructure contains all the necessary information for the script to work
type ParaStructure struct {
	Directories []ParaDirectory `toml:"directories"` // Needs to be an Exported field so the other libraries can detect it
}
