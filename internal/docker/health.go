package docker

import (
	"os/exec"
	"strings"
)

func IsContainerRunning(name string) bool {

	cmd := exec.Command(
		"docker",
		"inspect",
		"-f",
		"{{.State.Running}}",
		name,
	)

	out, err := cmd.Output()

	if err != nil {
		return false
	}

	return strings.TrimSpace(string(out)) == "true"
}