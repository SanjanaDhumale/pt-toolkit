package services

import (
	"fmt"

	"github.com/SanjanaDhumale/pt-toolkit/internal/installer"
)

func InstallStack(name string) error {

	switch name {

	case "protocol":
		return installer.InstallProtocolStack()

	case "monitoring":
		return installer.InstallMonitoringStack()

	default:
		return fmt.Errorf("unknown stack: %s", name)
	}
}