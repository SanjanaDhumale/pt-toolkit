package stack

import (
	"fmt"

	"github.com/SanjanaDhumale/pt-toolkit/internal/docker"
)

func DownMonitoring() error {

	fmt.Println()
	fmt.Println("Stopping Monitoring Stack")
	fmt.Println("----------------------------")

	containers := []string{
		"pt-grafana",
		"pt-prometheus",
		"pt-influxdb",
	}

	for _, c := range containers {

		if docker.IsContainerRunning(c) {

			docker.StopContainer(c)

			fmt.Printf("✓ %s stopped\n", c)

		} else {

			fmt.Printf("• %s already stopped\n", c)
		}
	}

	fmt.Println()
	fmt.Println("✓ Monitoring Stack Stopped")

	return nil
}