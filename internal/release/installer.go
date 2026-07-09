package release

import (
	"fmt"
	"os"
	"os/exec"
)

func BuildInstaller() error {

	fmt.Println("Step 5/5 : Building Windows Installer...")

	innoCompiler, err := FindISCC()
	if err != nil {
		fmt.Println("⚠", err)
		fmt.Println("Skipping installer generation.")
		fmt.Println()

		return nil
	}

	cmd := exec.Command(
		innoCompiler,
		"installer\\PTToolkit.iss",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Println("✓ Windows Installer Built")
	fmt.Println()

	return nil
}
