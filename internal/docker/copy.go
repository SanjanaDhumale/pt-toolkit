package docker

import (
	"os"
	"os/exec"
)

func CopyToContainer(container, source, destination string) error {

	cmd := exec.Command(
		"docker",
		"cp",
		source,
		container+":"+destination,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func CopyFromContainer(container, source, destination string) error {

	cmd := exec.Command(
		"docker",
		"cp",
		container+":"+source,
		destination,
	)

	return cmd.Run()
}
