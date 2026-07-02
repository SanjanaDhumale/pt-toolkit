package services

import (
	"fmt"

	dockerengine "github.com/SanjanaDhumale/pt-toolkit/internal/docker"
	"github.com/SanjanaDhumale/pt-toolkit/internal/registry"
)

func ShowImages() {

	fmt.Println("======================================")
	fmt.Println("Performance Engineering Toolkit Images")
	fmt.Println("======================================")
	fmt.Println()

	installed := 0
	missing := 0

	for _, service := range registry.AllServices() {

		if dockerengine.ImageExists(service.Image) {

			fmt.Printf("✓ %-15s Installed\n", service.Name)
			installed++

		} else {

			fmt.Printf("✗ %-15s Missing\n", service.Name)
			missing++

		}
	}

	fmt.Println()
	fmt.Println("--------------------------------------")
	fmt.Printf("Installed : %d\n", installed)
	fmt.Printf("Missing   : %d\n", missing)
	fmt.Println("======================================")
}