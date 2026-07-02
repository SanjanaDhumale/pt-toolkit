package services

import (
	"fmt"

	"github.com/SanjanaDhumale/pt-toolkit/internal/config"
	"github.com/SanjanaDhumale/pt-toolkit/internal/system"
)

func RunDoctor() {

	fmt.Println("==========================================")
	fmt.Println("PT Toolkit Environment Health Check")
	fmt.Println("==========================================")

	fmt.Printf("Toolkit Version : %s\n", config.AppConfig.Toolkit.Version)
	fmt.Printf("Docker Image    : %s\n\n", config.AppConfig.Docker.Image)

	checks := []struct {
		Name string
		Run  func() (string, error)
	}{
		{"Go", system.CheckGo},
		{"Docker", system.CheckDocker},
		{"Docker Compose", system.CheckDockerCompose},
		{"Java", system.CheckJava},
		{"Python", system.CheckPython},
		{"Git", system.CheckGit},
	}

	for _, check := range checks {

		version, err := check.Run()

		if err != nil {
			fmt.Printf("❌ %-18s Not Installed\n", check.Name)
			continue
		}

		fmt.Printf("✅ %-18s %s\n", check.Name, version)
	}

	fmt.Println("\n==========================================")
	fmt.Println("Environment Check Completed")
	fmt.Println("==========================================")

	fmt.Println()

	if system.IsDockerInstalled() {
		fmt.Println("✅ Docker Status : READY")
	} else {
		fmt.Println("❌ Docker Status : NOT INSTALLED")
		fmt.Println("Please install Docker Desktop before continuing.")
	}
}

