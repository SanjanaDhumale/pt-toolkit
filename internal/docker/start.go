package docker

import (
	"fmt"
	"os/exec"
)

func StartContainer(name, image string) error {

	fmt.Println("Starting:", name)

	cmd := exec.Command(
		"docker",
		"run",
		"-d",
		"--name",
		name,
		image,
	)

	out, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf(string(out))
	}

	return nil
}