package release

import (
	"fmt"
	"os"
	"os/exec"
)

func FindISCC() (string, error) {

	// 1. Check if ISCC is available in PATH
	if path, err := exec.LookPath("ISCC.exe"); err == nil {
		return path, nil
	}

	if path, err := exec.LookPath("ISCC"); err == nil {
		return path, nil
	}

	// 2. Check common installation locations
	locations := []string{
		`C:\Program Files (x86)\Inno Setup 6\ISCC.exe`,
		`C:\Program Files\Inno Setup 6\ISCC.exe`,
	}

	for _, path := range locations {
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}

	return "", fmt.Errorf("ISCC.exe not found.\nPlease install Inno Setup 6 or add ISCC.exe to PATH.")
}
