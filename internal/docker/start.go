package docker

import (
	"fmt"
	"os/exec"
)



func StartContainer(name, image string) error {

	if IsRunning(name) {

		fmt.Println("✓", name, "already running")

		return nil
	}

	fmt.Println("Starting:", name)

	cmd := exec.Command(
		"docker",
		"run",
		"-d",
		"--name",
		name,
		image,
	)

	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("%s", output)
	}

	fmt.Println("✓", name, "started")

	return nil
}