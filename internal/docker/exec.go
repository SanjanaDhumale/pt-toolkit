package docker

import (
	"os"
	"os/exec"
)

func Exec(container string, args ...string) error {

	command := []string{
		"exec",
		container,
	}

	command = append(command, args...)

	cmd := exec.Command("docker", command...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
