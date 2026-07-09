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

	fmt.Printf("%-12s %s\n", "Version:", config.AppConfig.Toolkit.Version)
	fmt.Printf("%-12s %s\n", "Build:", "Stable")
	fmt.Printf("%-12s %s\n", "Go:", runtime.Version())
	fmt.Printf("%-12s %s/%s\n", "Platform:", runtime.GOOS, runtime.GOARCH)
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
