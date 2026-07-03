package services

import (
	"fmt"

	"github.com/SanjanaDhumale/pt-toolkit/internal/docker"
	"github.com/SanjanaDhumale/pt-toolkit/internal/registry"
)

func StopService(name string) {

	serviceList, ok := registry.ToolRegistry[name]

	if !ok {
		fmt.Println("Unknown service:", name)
		return
	}

	fmt.Println("Stopping Services...")
	fmt.Println()

	for _, s := range serviceList {

		err := docker.StopContainer(s.Container)

		if err != nil {
			fmt.Println("✗", s.Name, err)
			continue
		}

		fmt.Println("✓", s.Name)
	}

	fmt.Println()
	fmt.Println("Environment Stopped")
}