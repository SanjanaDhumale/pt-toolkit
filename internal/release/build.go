package release

import (
	"fmt"
	"os"
	"os/exec"
)

func BuildExecutable() error {

	fmt.Println("Step 1/5 : Building executable...")

	cmd := exec.Command(
		"go",
		"build",
		"-o",
		"build/ptctl.exe",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Println("✓ Executable Built")
	fmt.Println()

	return nil
}
