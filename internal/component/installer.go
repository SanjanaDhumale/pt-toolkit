package component

import (
	"github.com/SanjanaDhumale/pt-toolkit/internal/ui"
)

func Install(c Component) Result {

	ui.Step("Checking " + c.Name())

	result := c.Check()

	if result.Status == Installed {

		ui.Success(result.Message)

		return result
	}

	ui.Warning(result.Message)

	ui.Step("Installing " + c.Name())

	install := c.Install()

	if install.Status == Installed {

		ui.Success(install.Message)

	} else {

		ui.Error(install.Message)
	}

	return install
}