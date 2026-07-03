package installer

import "fmt"

func Install() {

	fmt.Println("========================================")
	fmt.Println(" PT Toolkit Installer")
	fmt.Println("========================================")

	CheckDocker()

	InstallImages()

	fmt.Println()

	fmt.Println("Installation Completed")
}