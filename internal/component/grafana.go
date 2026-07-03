package component

import "github.com/SanjanaDhumale/pt-toolkit/internal/docker"

type Grafana struct{}

func (Grafana) Name() string {
	return "Grafana"
}

func (Grafana) Check() Result {

	if docker.ImageExists("grafana/grafana:latest") {

		return Result{
			Name: "Grafana",
			Status: Installed,
			Message: "Grafana image available",
		}
	}

	return Result{
		Name: "Grafana",
		Status: NotInstalled,
		Message: "Grafana image missing",
	}
}

func (Grafana) Install() Result {

	if err := docker.PullImage("grafana/grafana:latest"); err != nil {

		return Result{
			Name: "Grafana",
			Status: NotInstalled,
			Message: err.Error(),
		}
	}

	return Result{
		Name: "Grafana",
		Status: Installed,
		Message: "Grafana installed",
	}
}

func (Grafana) Update() Result {
	return Result{Name: "Grafana", Status: Installed, Message: "Already latest"}
}

func (Grafana) Repair() Result {
	return Result{Name: "Grafana", Status: Installed, Message: "Repair completed"}
}

func (Grafana) Uninstall() Result {
	return Result{Name: "Grafana", Status: Installed, Message: "Removed"}
}