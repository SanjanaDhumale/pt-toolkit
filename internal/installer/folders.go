package installer

import (
	"fmt"
	"os"
)

var folders = []string{
	"config",
	"reports",
	"logs",
	"scripts",
	"plugins",
	"templates",
	"workspace",
}

func CreateFolders() {

	fmt.Println()
	fmt.Println("Creating Toolkit Folders")

	for _, folder := range folders {

		err := os.MkdirAll(folder, 0755)

		if err != nil {
			fmt.Println("✗", folder)
			continue
		}

		fmt.Println("✓", folder)
	}
}