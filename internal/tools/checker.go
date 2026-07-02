package tools

import (
	"fmt"
	"os/exec"

	dockerengine "github.com/SanjanaDhumale/pt-toolkit/internal/docker"
)

type Tool struct {
	Name    string
	Command string
	Image   string
	IsImage bool
}

var toolList = []Tool{
	{"Java", "java", "", false},
	{"Python", "python", "", false},
	{"Git", "git", "", false},
	{"Docker", "docker", "", false},
	{"Docker Compose", "docker", "", false},

	{"JMeter", "", "pt-jmeter:v1", true},
	{"Grafana", "", "grafana/grafana:latest", true},
	{"Prometheus", "", "prom/prometheus:latest", true},
	{"InfluxDB", "", "influxdb:2.7", true},
}

func CheckAll() {

	fmt.Println("==========================================")
	fmt.Println(" Performance Engineering Tool Manager")
	fmt.Println("==========================================")
	fmt.Println()

	installed := 0
	missing := 0

	for _, tool := range toolList {

		ok := false

		if tool.IsImage {

			ok = dockerengine.ImageExists(tool.Image)

		} else {

			_, err := exec.LookPath(tool.Command)
			ok = err == nil
		}

		if ok {

			fmt.Printf("✓ %-18s Installed\n", tool.Name)
			installed++

		} else {

			fmt.Printf("✗ %-18s Missing\n", tool.Name)
			missing++
		}
	}

	fmt.Println()
	fmt.Println("------------------------------------------")
	fmt.Printf("Installed : %d\n", installed)
	fmt.Printf("Missing   : %d\n", missing)
	fmt.Println("==========================================")
}