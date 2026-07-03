package component

import (
	"github.com/SanjanaDhumale/pt-toolkit/internal/ui"
)

func Install(c Component) error {

	ui.Step("Checking " + c.Name())

	if err := c.Check(); err == nil {

		ui.Success(c.Name() + " already installed")

		return nil
	}

	ui.Step("Installing " + c.Name())

	if err := c.Install(); err != nil {

		ui.Error(err.Error())

		return err
	}

	ui.Success(c.Name() + " installed")

	return nil
}