package docker

import (
	"fmt"
	"os"
	"os/exec"
)

func ListImages() {

	cmd := exec.Command("docker", "images")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to list Docker images:", err)
	}
}