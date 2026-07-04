package stack

import (
	"fmt"

	"github.com/SanjanaDhumale/pt-toolkit/internal/docker"
)

func StatusMonitoring() {

	fmt.Println()
	fmt.Println("Monitoring Stack")
	fmt.Println("---------------------------")

	containers := []string{
		"pt-grafana",
		"pt-prometheus",
		"pt-influxdb",
	}

	for _, c := range containers {

		if docker.IsContainerRunning(c) {
			fmt.Printf("✓ %-20s Running\n", c)
		} else {
			fmt.Printf("✗ %-20s Stopped\n", c)
		}
	}
}