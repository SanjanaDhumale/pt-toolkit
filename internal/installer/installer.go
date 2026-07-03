package installer

import "fmt"

func Install() {

	CheckDocker()

	InstallImages()

	CreateFolders()

	CreateConfig()

	fmt.Println()

	fmt.Println("✓ Installation Completed")
}