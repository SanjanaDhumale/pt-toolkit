package engine

import (
	"github.com/SanjanaDhumale/pt-toolkit/internal/component"
	"github.com/SanjanaDhumale/pt-toolkit/internal/stack"
	"github.com/SanjanaDhumale/pt-toolkit/internal/ui"
	"github.com/SanjanaDhumale/pt-toolkit/internal/docker"
)

func Install(stackName string) error {

	s, err := stack.Find(stackName)
	if err != nil {
		return err
	}

	ui.Header("Installing " + s.Name)

	ui.Step("Creating Docker Network")

	if err := docker.CreateNetwork(s.Network); err != nil {
		return err
	}

	ui.Success(s.Network + " ready")

	for _, c := range s.Components {

		result := c.Check()

		switch result.Status {

		case component.Installed:

			ui.Success(result.Message)

		case component.NotInstalled:

			ui.Warning(result.Message)

			install := c.Install()

			if install.Status == component.Installed {

				ui.Success(install.Message)

			} else {

				ui.Error(install.Message)
			}
		}
	}

	ui.Success(s.Name + " Installed")

	return nil
}