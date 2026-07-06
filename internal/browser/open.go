package browser

import (
	"os/exec"
	"runtime"
)

func Open(path string) error {

	switch runtime.GOOS {

	case "windows":
		return exec.Command(
			"cmd",
			"/c",
			"start",
			"",
			path,
		).Run()

	case "darwin":
		return exec.Command(
			"open",
			path,
		).Run()

	default:
		return exec.Command(
			"xdg-open",
			path,
		).Run()
	}
}
