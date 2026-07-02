package services

import (
	"fmt"

	"github.com/SanjanaDhumale/pt-toolkit/internal/registry"
)

func ListServices() {

	fmt.Println("====================================")
	fmt.Println("Available Performance Services")
	fmt.Println("====================================")

	services := registry.AllServices()

	for i, service := range services {

		fmt.Printf("%d. %-15s %s\n",
			i+1,
			service.Name,
			service.Description,
		)
	}
}