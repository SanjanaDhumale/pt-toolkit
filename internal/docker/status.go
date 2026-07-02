package docker

import (
	"os/exec"
	"strings"
)

func IsRunning(name string) bool {

	cmd := exec.Command(
		"docker",
		"ps",
		"--format",
		"{{.Names}}",
	)

	out, err := cmd.Output()

	if err != nil {
		return false
	}

	return strings.Contains(string(out), name)

}