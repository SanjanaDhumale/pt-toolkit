package docker

import "os/exec"

func StartExistingContainer(name string) error {

	cmd := exec.Command(
		"docker",
		"start",
		name,
	)

	return cmd.Run()
}