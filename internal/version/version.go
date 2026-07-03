package version

import (
	"fmt"
	"runtime"

	"github.com/SanjanaDhumale/pt-toolkit/internal/config"
	"github.com/SanjanaDhumale/pt-toolkit/internal/registry"
)

func Show() {

	fmt.Println("====================================")
	fmt.Println(config.AppConfig.Toolkit.Name)
	fmt.Println("====================================")
	fmt.Println()

	fmt.Println("Version :", config.AppConfig.Toolkit.Version)
	fmt.Println("Go      :", runtime.Version())
	fmt.Println("OS      :", runtime.GOOS)
	fmt.Println()

	fmt.Println("Supported Services")

	seen := make(map[string]bool)

	for _, services := range registry.ToolRegistry {
		for _, s := range services {

			if seen[s.Name] {
				continue
			}

			seen[s.Name] = true

			fmt.Println("✓", s.Name)
		}
	}

	fmt.Println()

	fmt.Println("====================================")
}