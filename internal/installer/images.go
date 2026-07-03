package installer

import (
	"fmt"

	dockerengine "github.com/SanjanaDhumale/pt-toolkit/internal/docker"
)

var requiredImages = []string{
	"pt-jmeter:v1",
	"grafana/grafana:latest",
	"prom/prometheus:latest",
	"influxdb:2.7",
	"selenium/standalone-chrome:latest",
}

func InstallImages() {

	fmt.Println()
	fmt.Println("Checking Docker Images")
	fmt.Println()

	for _, image := range requiredImages {

		if dockerengine.ImageExists(image) {

			fmt.Println("✓", image)

		} else {

			fmt.Println("⬇ Downloading", image)

			dockerengine.PullImage(image)
		}
	}
}