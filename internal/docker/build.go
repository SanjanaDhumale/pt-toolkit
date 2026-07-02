package docker

import (
	"fmt"
	"os/exec"
)

func BuildImage(tag, dockerfilePath string) error {

	fmt.Println("Building image:", tag)

	cmd := exec.Command(
		"docker",
		"build",
		"-t",
		tag,
		dockerfilePath,
	)

	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("%s", output)
	}

	return nil
}