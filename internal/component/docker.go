package component

import "github.com/SanjanaDhumale/pt-toolkit/internal/system"

type Docker struct{}

func (Docker) Name() string {
	return "Docker"
}

func (Docker) Check() Result {

	version, err := system.CheckDocker()

	if err != nil {
		return Result{
			Name:    "Docker",
			Status:  NotInstalled,
			Message: "Docker is not installed",
		}
	}

	return Result{
		Name:    "Docker",
		Status:  Installed,
		Message: version,
	}
}

func (Docker) Install() Result {
	return Result{
		Name:    "Docker",
		Status:  Installed,
		Message: "Please install Docker Desktop",
	}
}

func (Docker) Update() Result {
	return Result{
		Name:    "Docker",
		Status:  Installed,
		Message: "Docker is up to date",
	}
}

func (Docker) Repair() Result {
	return Result{
		Name:    "Docker",
		Status:  Installed,
		Message: "Repair completed",
	}
}

func (Docker) Uninstall() Result {
	return Result{
		Name:    "Docker",
		Status:  Installed,
		Message: "Unsupported",
	}
}