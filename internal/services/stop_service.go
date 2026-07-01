package services

import (
	"fmt"
	"os"
	"os/exec"
	"github.com/SanjanaDhumale/pt-toolkit/internal/config"
)

func RunStop() {

	fmt.Println("===================================")
	fmt.Println("Stopping PT Toolkit...")
	fmt.Println("===================================")

	cmd := exec.Command(
		"docker",
		"compose",
		"-f",
		config.AppConfig.Docker.ComposeFile,
		"down",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()

	if err != nil {
		fmt.Println("Failed to stop PT Toolkit:", err)
		return
	}

	fmt.Println()
	fmt.Println("✅ PT Toolkit stopped successfully.")
}