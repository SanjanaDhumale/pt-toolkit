package services

import (
	"fmt"
	"os"
	"os/exec"
	"github.com/SanjanaDhumale/pt-toolkit/internal/config"
)

func RunInstall() {

	fmt.Println("===================================")
	fmt.Println("Building PT Toolkit Docker Image...")
	fmt.Println("===================================")

	cmd := exec.Command(
		"docker",
		"build",
		"-f",
		config.AppConfig.Docker.Dockerfile,
		"-t",
		config.AppConfig.Docker.Image,
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