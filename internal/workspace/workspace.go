package workspace

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateWorkspace(projectName string) error {

	folders := []string{
		"configs",
		"scripts",
		"testdata",
		"reports",
		"logs",
		"results",
		"templates",
		"plugins",
	}

	for _, folder := range folders {

		path := filepath.Join(projectName, folder)

		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}

	
	fmt.Println("Workspace Created Successfully")

	config := `project:
	name: "` + projectName + `"

	jmeter:
	enabled: true

	k6:
	enabled: false

	selenium:
	enabled: false

	monitoring:
	enabled: false
	`

	configPath := filepath.Join(projectName, "configs", "toolkit.yaml")

	err := os.WriteFile(configPath, []byte(config), 0644)

	if err != nil {
		return err
	}
	
	fmt.Println("Project :", projectName)

	return nil
}

