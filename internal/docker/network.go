package docker

import (
	"os/exec"
)

func NetworkExists(name string) bool {

	cmd := exec.Command(
		"docker",
		"network",
		"inspect",
		name,
	)

	return cmd.Run() == nil
}

func CreateNetwork(name string) error {

	if NetworkExists(name) {
		return nil
	}

	cmd := exec.Command(
		"docker",
		"network",
		"create",
		name,
	)

	return cmd.Run()
}