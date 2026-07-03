package component

import (
	"github.com/SanjanaDhumale/pt-toolkit/internal/system"
)

type Java struct{}

func (Java) Name() string {
	return "Java"
}

func (Java) Check() Result {

	version, err := system.CheckJava()

	if err != nil {

		return Result{
			Name:    "Java",
			Status:  NotInstalled,
			Message: "Java not installed",
		}
	}

	return Result{
		Name:    "Java",
		Status:  Installed,
		Message: version,
	}
}

func (Java) Install() Result {

	return Result{
		Name:    "Java",
		Status:  Installed,
		Message: "Please install Java 21",
	}
}

func (Java) Update() Result {
	return Result{Name: "Java", Message: "Already Latest"}
}

func (Java) Repair() Result {
	return Result{Name: "Java", Message: "Repair Completed"}
}

func (Java) Uninstall() Result {
	return Result{Name: "Java", Message: "Unsupported"}
}