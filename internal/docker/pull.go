package docker

import (
	"fmt"
	"os/exec"
)

func PullImage(image string) error {

	fmt.Println("Pulling image:", image)

	cmd := exec.Command("docker", "pull", image)

	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("%s", output)
	}

	return nil
}