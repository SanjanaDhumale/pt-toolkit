package workspace

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CurrentWorkspace() (string, error) {

	file := filepath.Join("config", "current-workspace.txt")

	data, err := os.ReadFile(file)

	if err != nil {
		return "", fmt.Errorf("no active workspace")
	}

	return strings.TrimSpace(string(data)), nil
}
