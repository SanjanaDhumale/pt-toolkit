package services

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"github.com/SanjanaDhumale/pt-toolkit/internal/config"
	"github.com/SanjanaDhumale/pt-toolkit/internal/execution"
)

func ImageExists() bool {

	cmd := exec.Command(
		"docker",
		"image",
		"inspect",
		config.AppConfig.Docker.Image,
	)

	err := cmd.Run()

	return err == nil
}

// Build Docker Image
func RunInstall() {
	
	fmt.Println("Installing PT Toolkit...")

	if ImageExists() {
		fmt.Println("✅ Docker image already exists.")
		fmt.Println("Skipping image build.")
		return
	}

	fmt.Println("Docker image not found.")
	fmt.Println("Building image...")

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

	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Installation Failed:", err)
		return
	}

	fmt.Println("\n✅ PT Toolkit installed successfully.")
}


func ContainerExists(containerName string) bool {

	cmd := exec.Command(
		"docker",
		"ps",
		"-a",
		"--filter",
		"name="+containerName,
		"--format",
		"{{.Names}}",
	)

	output, err := cmd.Output()

	if err != nil {
		return false
	}

	return strings.Contains(string(output), containerName)
}

// Start Docker Environment
func RunStart(tool string) {

	err := execution.SelectTool(tool)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("===================================")
	fmt.Println("Starting PT Toolkit...")
	fmt.Println("===================================")

	err = StartRequiredContainers(tool)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()
	fmt.Println("✅ Environment Ready")
}


// Stop Docker Environment
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

	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Failed to stop PT Toolkit:", err)
		return
	}

	fmt.Println("\n✅ PT Toolkit stopped successfully.")
}

// Show Container Status
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

	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Error:", err)
	}
}