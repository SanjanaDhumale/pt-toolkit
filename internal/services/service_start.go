package services

import (
	"fmt"

	"github.com/SanjanaDhumale/pt-toolkit/internal/docker"
	"github.com/SanjanaDhumale/pt-toolkit/internal/registry"
)

func StartService(name string) {

	serviceList, ok := registry.ToolRegistry[name]

	if !ok {

		fmt.Println("Unknown service:", name)
		return
	}

	fmt.Println("Starting Services...")
	fmt.Println()

	for _, s := range serviceList {

		err := docker.StartContainer(
			s.Container,
			s.Image,
		)

		if err != nil {

			fmt.Println("✗", s.Name, err)
			continue
		}

		fmt.Println("✓", s.Name)
	}

	fmt.Println()
	fmt.Println("Environment Ready")
}