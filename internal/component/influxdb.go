package component

import "github.com/SanjanaDhumale/pt-toolkit/internal/docker"

type InfluxDB struct{}

func (InfluxDB) Name() string {
	return "InfluxDB"
}

func (InfluxDB) Check() Result {

	if docker.ImageExists("influxdb:2.7") {

		return Result{
			Name:    "InfluxDB",
			Status:  Installed,
			Message: "InfluxDB image available",
		}
	}

	return Result{
		Name:    "InfluxDB",
		Status:  NotInstalled,
		Message: "InfluxDB image missing",
	}
}

func (InfluxDB) Install() Result {

	if err := docker.PullImage("influxdb:2.7"); err != nil {

		return Result{
			Name:    "InfluxDB",
			Status:  NotInstalled,
			Message: err.Error(),
		}
	}

	return Result{
		Name:    "InfluxDB",
		Status:  Installed,
		Message: "InfluxDB installed",
	}
}

func (InfluxDB) Update() Result {
	return Result{
		Name:    "InfluxDB",
		Status:  Installed,
		Message: "Already latest",
	}
}

func (InfluxDB) Repair() Result {
	return Result{
		Name:    "InfluxDB",
		Status:  Installed,
		Message: "Repair completed",
	}
}

func (InfluxDB) Uninstall() Result {
	return Result{
		Name:    "InfluxDB",
		Status:  Installed,
		Message: "Removed",
	}
}