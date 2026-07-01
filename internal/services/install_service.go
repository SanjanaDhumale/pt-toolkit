package services

import (
	"fmt"
	"os"
	"os/exec"
)

func RunInstall() {

	fmt.Println("===================================")
	fmt.Println("Building PT Toolkit Docker Image...")
	fmt.Println("===================================")

	cmd := exec.Command(
		"docker",
		"build",
		"-f",
		"docker/Dockerfile",
		"-t",
		"pt-toolkit:v1",
		".",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()

	if err != nil {
		fmt.Println("Installation Failed:", err)
		return
	}

	fmt.Println()
	fmt.Println("✅ PT Toolkit installed successfully.")
}