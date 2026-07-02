package environment

import (
	"fmt"
	"os"
	"path/filepath"
)

var folders = []string{
	"PTToolkit/configs",
	"PTToolkit/logs",
	"PTToolkit/reports",
	"PTToolkit/plugins",
	"PTToolkit/data",
	"PTToolkit/templates",
	"PTToolkit/docker",
}

func Setup() {

	fmt.Println()
	fmt.Println("Creating Toolkit Environment")

	for _, folder := range folders {

		err := os.MkdirAll(filepath.Join(".", folder), os.ModePerm)

		if err != nil {

			fmt.Println("✗", folder)

			continue
		}

		fmt.Println("✓", folder)
	}
}