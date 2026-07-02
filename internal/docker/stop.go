package docker

import (
	"os/exec"
)

func StopContainer(name string) {

	exec.Command(
		"docker",
		"stop",
		name,
	).Run()

	exec.Command(
		"docker",
		"rm",
		name,
	).Run()

}