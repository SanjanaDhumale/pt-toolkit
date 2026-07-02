package services

import (
	"fmt"

	"github.com/SanjanaDhumale/pt-toolkit/internal/docker"
	"github.com/SanjanaDhumale/pt-toolkit/internal/registry"
)

func ServiceStatus() {

	fmt.Println("=========================================")
	fmt.Println(" Performance Services Status")
	fmt.Println("=========================================")

	all := registry.AllServices()

	for _, service := range all {

		if docker.IsRunning(service.Container) {

			fmt.Printf("✓ %-15s Running\n", service.Name)

		} else {

			fmt.Printf("✗ %-15s Stopped\n", service.Name)

		}
	}

	fmt.Println("=========================================")
}