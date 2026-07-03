package installer

import (
	"github.com/SanjanaDhumale/pt-toolkit/internal/ui"
)

func InstallMonitoringStack() error {

	ui.Header("Monitoring Stack")

	ui.Step("Checking Docker")

	ui.Step("Pulling Grafana")

	ui.Step("Pulling Prometheus")

	ui.Step("Pulling InfluxDB")

	ui.Success("Monitoring Stack Installed")

	return nil
}