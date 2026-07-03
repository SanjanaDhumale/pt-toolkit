package installer

import (
	"fmt"
	"os"
)

const defaultConfig = `toolkit:
  name: PT Toolkit
  version: 0.1.0

docker:
  image: pt-toolkit:v1

plugins:
  jmeter: true
  k6: true
  selenium: true
`

func CreateConfig() {

	file := "config/toolkit.yaml"

	if _, err := os.Stat(file); err == nil {
		fmt.Println("✓ toolkit.yaml already exists")
		return
	}

	err := os.WriteFile(file, []byte(defaultConfig), 0644)

	if err != nil {
		fmt.Println("✗ Failed creating toolkit.yaml")
		return
	}

	fmt.Println("✓ toolkit.yaml created")
}