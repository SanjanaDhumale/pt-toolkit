package engine

import (
	"github.com/SanjanaDhumale/pt-toolkit/internal/component"
	"github.com/SanjanaDhumale/pt-toolkit/internal/stack"
	"github.com/SanjanaDhumale/pt-toolkit/internal/ui"
)

func Install(stackName string) error {

	s, err := stack.Find(stackName)
	if err != nil {
		return err
	}

	ui.Header("Installing " + s.Name)

	// TODO:
	// Later, s.Components will contain actual Component implementations.
	// For now, we'll manually execute the Java component.
	_ = s

	java := component.Java{}

	result := java.Check()

	switch result.Status {

	case component.Installed:
		ui.Success(result.Message)

	case component.NotInstalled:
		ui.Warning(result.Message)

		install := java.Install()
		ui.Info(install.Message)
	}

	ui.Success(s.Name + " Installed")

	return nil
}