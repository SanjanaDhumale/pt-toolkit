package services

import (
	"fmt"
	"os"
	"os/exec"
)

func RunStatus() {

	fmt.Println("===================================")
	fmt.Println("PT Toolkit Status")
	fmt.Println("===================================")

	cmd := exec.Command(
		"docker",
		"ps",
		"-a",
		"--filter",
		"name=pt-toolkit",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		fmt.Println("Error:", err)
	}
}