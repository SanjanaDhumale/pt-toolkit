package docker

import (
	"os"
	"os/exec"
)

func GenerateHTMLReport(container string) error {

	// Delete previous dashboard
	exec.Command(
		"docker",
		"exec",
		container,
		"rm",
		"-rf",
		"/tmp/dashboard",
	).Run()

	cmd := exec.Command(
		"docker",
		"exec",
		container,
		"jmeter",
		"-g",
		"/tmp/result.jtl",
		"-o",
		"/tmp/dashboard",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
