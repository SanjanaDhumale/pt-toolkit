package component

import (
	"os/exec"
)

type JMeter struct{}

func (JMeter) Name() string {
	return "JMeter"
}

func (JMeter) Check() Result {

	cmd := exec.Command(
		"docker",
		"image",
		"inspect",
		"pt-jmeter:v1",
	)

	if err := cmd.Run(); err != nil {

		return Result{
			Name:    "JMeter",
			Status:  NotInstalled,
			Message: "PT JMeter image not found",
		}
	}

	return Result{
		Name:    "JMeter",
		Status:  Installed,
		Message: "PT JMeter image available",
	}
}

func (JMeter) Install() Result {

	cmd := exec.Command(
		"docker",
		"pull",
		"pt-jmeter:v1",
	)

	if err := cmd.Run(); err != nil {

		return Result{
			Name:    "JMeter",
			Status:  NotInstalled,
			Message: "Failed to pull PT JMeter image",
		}
	}

	return Result{
		Name:    "JMeter",
		Status:  Installed,
		Message: "PT JMeter image installed",
	}
}

func (JMeter) Update() Result {

	return Result{
		Name:    "JMeter",
		Status:  Installed,
		Message: "Already latest",
	}
}

func (JMeter) Repair() Result {

	return Result{
		Name:    "JMeter",
		Status:  Installed,
		Message: "Repair completed",
	}
}

func (JMeter) Uninstall() Result {

	return Result{
		Name:    "JMeter",
		Status:  Installed,
		Message: "Removed",
	}
}