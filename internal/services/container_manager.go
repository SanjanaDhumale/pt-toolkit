package services

import (
	"fmt"

	dockerengine "github.com/SanjanaDhumale/pt-toolkit/internal/docker"
	"github.com/SanjanaDhumale/pt-toolkit/internal/registry"
)



func StartRequiredContainers(tool string) error {

	services, err := registry.Lookup(tool)

	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("Starting Required Containers")
	fmt.Println()

	for _, service := range services {

		if dockerengine.IsRunning(service.Container) {

			fmt.Printf("✓ %s already running\n", service.Name)
			continue
		}

		err := dockerengine.StartContainer(
			service.Container,
			service.Image,
		)

		if err != nil {

			fmt.Printf("✗ %s : %v\n", service.Name, err)
			continue
		}

		fmt.Printf("✓ %s started\n", service.Name)

	}

	return nil
}