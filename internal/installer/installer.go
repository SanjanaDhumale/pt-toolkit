package installer

import "fmt"

func Install() {

	fmt.Println("========================================")
	fmt.Println("PT Toolkit Installer")
	fmt.Println("========================================")

	CheckDocker()

	InstallImages()

	CreateFolders()

	fmt.Println()

	fmt.Println("✓ Installation Completed")
}