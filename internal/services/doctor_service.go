package services

import (
	"fmt"
	"os"
	"os/exec"
	"github.com/SanjanaDhumale/pt-toolkit/internal/system"
	"github.com/SanjanaDhumale/pt-toolkit/internal/config"
	
)

func RunDoctor() {

	fmt.Println("======================================")
	fmt.Println(" PT Toolkit Environment Health Check")
	fmt.Println("======================================")

	goVersion, _ := system.CheckGo()
	fmt.Println(goVersion)

	dockerVersion, _ := system.CheckDocker()
	fmt.Println(dockerVersion)

	composeVersion, _ := system.CheckDockerCompose()
	fmt.Println(composeVersion)

	javaVersion, _ := system.CheckJava()
	fmt.Println(javaVersion)

	pythonVersion, _ := system.CheckPython()
	fmt.Println(pythonVersion)

	gitVersion, _ := system.CheckGit()
	fmt.Println(gitVersion)
}

func RunStart() {

	fmt.Println("====================================")
	fmt.Println("Starting PT Toolkit...")
	fmt.Println("====================================")

	cmd := exec.Command(
    "docker",
    "compose",
    "-f",
    config.AppConfig.Docker.ComposeFile,
    "up",
)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()

	if err != nil {
		fmt.Println("Error:", err)
	}
}