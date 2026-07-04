package execution

import (
	"fmt"

	"github.com/SanjanaDhumale/pt-toolkit/internal/config"
	"github.com/SanjanaDhumale/pt-toolkit/internal/registry"
)

func SelectTool(tool string) error {

	switch tool {

	case "jmeter":

		if !config.AppConfig.Tools.JMeter.Enabled {
			return fmt.Errorf("JMeter disabled")
		}

	case "k6":

		if !config.AppConfig.Tools.K6.Enabled {
			return fmt.Errorf("K6 disabled")
		}

	case "selenium":

		if !config.AppConfig.Tools.Selenium.Enabled {
			return fmt.Errorf("Selenium disabled")
		}

	default:
		return fmt.Errorf("unknown tool")
	}

	services, err := registry.Lookup(tool)

	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("Checking Required Services")
	fmt.Println()

	for _, service := range services {

		fmt.Printf("✓ %-15s %s\n",
			service.Name,
			service.Description)

	}

	fmt.Println()

	return nil
}