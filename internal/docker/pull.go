package docker

import (
	"fmt"
	"os"
	"os/exec"
)

func PullImage(image string) {

	cmd := exec.Command(
		"docker",
		"pull",
		image,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {

		fmt.Println("❌ Failed:", image)

		return
	}

	fmt.Println("✓ Downloaded:", image)
}