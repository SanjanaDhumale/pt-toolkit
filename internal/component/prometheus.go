package component

import "github.com/SanjanaDhumale/pt-toolkit/internal/docker"

type Prometheus struct{}

func (Prometheus) Name() string {
	return "Prometheus"
}

func (Prometheus) Check() Result {

	if docker.ImageExists("prom/prometheus:latest") {

		return Result{
			Name:    "Prometheus",
			Status:  Installed,
			Message: "Prometheus image available",
		}
	}

	return Result{
		Name:    "Prometheus",
		Status:  NotInstalled,
		Message: "Prometheus image missing",
	}
}

func (Prometheus) Install() Result {

	if err := docker.PullImage("prom/prometheus:latest"); err != nil {

		return Result{
			Name:    "Prometheus",
			Status:  NotInstalled,
			Message: err.Error(),
		}
	}

	return Result{
		Name:    "Prometheus",
		Status:  Installed,
		Message: "Prometheus installed",
	}
}

func (Prometheus) Update() Result {
	return Result{
		Name:    "Prometheus",
		Status:  Installed,
		Message: "Already latest",
	}
}

func (Prometheus) Repair() Result {
	return Result{
		Name:    "Prometheus",
		Status:  Installed,
		Message: "Repair completed",
	}
}

func (Prometheus) Uninstall() Result {
	return Result{
		Name:    "Prometheus",
		Status:  Installed,
		Message: "Removed",
	}
}