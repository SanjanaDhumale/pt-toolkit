package installer

import (
	"github.com/SanjanaDhumale/pt-toolkit/internal/ui"
)

func InstallProtocolStack() error {

	ui.Header("Protocol Testing Stack")

	ui.Step("Checking Java")
	// TODO: Check Java

	ui.Step("Checking Docker")
	// TODO: Check Docker

	ui.Step("Pulling PT JMeter Image")
	// TODO: Pull Image

	ui.Step("Installing JMeter Plugins")
	// TODO: Install Plugins

	ui.Step("Creating Script Templates")
	// TODO: Copy Templates

	ui.Success("Protocol Stack Installed")

	return nil
}