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

func CheckImages() {

	for _, image := range requiredImages {

		if dockerengine.ImageExists(image) {

			fmt.Println("✓", image)

		} else {

			fmt.Println("⬇ Pulling", image)

			dockerengine.PullImage(image)
		}
	}
}