package docker

import (
	"fmt"
	"os"
	"os/exec"
)

func PullImage(image string) error {

	cmd := exec.Command(
		"docker",
		"pull",
		image,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Println("✓", image, "downloaded")

	return nil
}