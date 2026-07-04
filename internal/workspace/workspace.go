package workspace

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CreateWorkspace(projectName string) error {

	// Workspace root
	root := filepath.Join("workspace", projectName)

	// Project folders
	folders := []string{
		"config",

		"scripts",
		"scripts/jmeter",
		"scripts/k6",
		"scripts/selenium",

		"data",
		"data/csv",
		"data/json",
		"data/payloads",

		"reports",
		"logs",
		"results",

		"plugins",
		"templates",
	}

	for _, folder := range folders {

		path := filepath.Join(root, folder)

		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// -------------------------
	// Project Configuration
	// -------------------------

	config := `project:
  name: "` + projectName + `"

tools:
  jmeter: true
  k6: false
  selenium: false

monitoring:
  grafana: false
  prometheus: false
  influxdb: false
`

	configPath := filepath.Join(root, "config", "config.yaml")

	err := os.WriteFile(configPath, []byte(config), 0644)

	if err != nil {
		return err
	}

	// -------------------------
	// README
	// -------------------------

	readme := `# ` + projectName + `

Created by PT Toolkit

This workspace contains:

- JMeter Scripts
- K6 Scripts
- Selenium Scripts
- Test Data
- Reports
- Logs
- Plugins
- Templates
`

	readmePath := filepath.Join(root, "README.md")

	err = os.WriteFile(readmePath, []byte(readme), 0644)

	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("======================================")
	fmt.Println(" Workspace Created Successfully")
	fmt.Println("======================================")
	fmt.Println()

	fmt.Println("Project :", projectName)
	fmt.Println("Location:", root)

	return nil
}


func ListWorkspaces() {

	root := "workspace"

	entries, err := os.ReadDir(root)

	if err != nil {

		fmt.Println("No workspaces found.")
		return
	}

	fmt.Println()
	fmt.Println("Available Workspaces")
	fmt.Println("====================")
	fmt.Println()

	for _, entry := range entries {

		if entry.IsDir() {

			fmt.Println("📁", entry.Name())

		}
	}

	fmt.Println()
}


func UseWorkspace(name string) error {

	root := filepath.Join("workspace", name)

	// Check if workspace exists
	if _, err := os.Stat(root); os.IsNotExist(err) {
		return fmt.Errorf("workspace '%s' not found", name)
	}

	err := os.MkdirAll("config", 0755)
	if err != nil {
		return err
	}

	file := filepath.Join("config", "current-workspace.txt")

	err = os.WriteFile(file, []byte(name), 0644)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("Current Workspace :", name)
	fmt.Println()

	return nil
}


func countFiles(folder string) int {

	entries, err := os.ReadDir(folder)

	if err != nil {
		return 0
	}

	count := 0

	for _, e := range entries {

		if !e.IsDir() {
			count++
		}
	}

	return count
}


func Info() error {

	data, err := os.ReadFile(filepath.Join("config", "current-workspace.txt"))

	if err != nil {
		return fmt.Errorf("no workspace selected")
	}

	name := strings.TrimSpace(string(data))

	root := filepath.Join("workspace", name)

	fmt.Println()
	fmt.Println("====================================")
	fmt.Println(" Current Workspace")
	fmt.Println("====================================")
	fmt.Println()

	fmt.Println("Name      :", name)
	fmt.Println("Location  :", root)
	fmt.Println()

	fmt.Println("Scripts   :", countFiles(filepath.Join(root, "scripts", "jmeter")))
	fmt.Println("Reports   :", countFiles(filepath.Join(root, "reports")))
	fmt.Println("Logs      :", countFiles(filepath.Join(root, "logs")))
	fmt.Println("Data Files:", countFiles(filepath.Join(root, "data", "csv")))

	fmt.Println()

	return nil
}