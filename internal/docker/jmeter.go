package docker

import (
	"os"
	"os/exec"
)

func RunJMeter(container, script, result string) error {

	cmd := exec.Command(
		"docker",
		"exec",
		container,
		"jmeter",
		"-n",
		"-t", script,
		"-l", result,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
