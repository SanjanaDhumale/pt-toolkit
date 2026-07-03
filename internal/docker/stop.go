package docker

import (
	"fmt"
	"os/exec"
)

func StopContainer(name string) error {

	if !IsRunning(name) {
		fmt.Println("✓", name, "already stopped")
		return nil
	}

	cmd := exec.Command(
		"docker",
		"stop",
		name,
	)

	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("%s", output)
	}

	cmd = exec.Command(
		"docker",
		"rm",
		name,
	)

	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("%s", output)
	}

	return nil
}