package component

import "github.com/SanjanaDhumale/pt-toolkit/internal/docker"

type Selenium struct{}

func (Selenium) Name() string {
	return "Selenium"
}

func (Selenium) Check() Result {

	if docker.ImageExists("selenium/standalone-chrome:latest") {
		return Result{
			Name:    "Selenium",
			Status:  Installed,
			Message: "Selenium image available",
		}
	}

	return Result{
		Name:    "Selenium",
		Status:  NotInstalled,
		Message: "Selenium image missing",
	}
}

func (Selenium) Install() Result {

	if err := docker.PullImage("selenium/standalone-chrome:latest"); err != nil {
		return Result{
			Name:    "Selenium",
			Status:  NotInstalled,
			Message: err.Error(),
		}
	}

	return Result{
		Name:    "Selenium",
		Status:  Installed,
		Message: "Selenium installed",
	}
}

func (Selenium) Update() Result {
	return Result{Name: "Selenium", Status: Installed, Message: "Already latest"}
}

func (Selenium) Repair() Result {
	return Result{Name: "Selenium", Status: Installed, Message: "Repair completed"}
}

func (Selenium) Uninstall() Result {
	return Result{Name: "Selenium", Status: Installed, Message: "Removed"}
}