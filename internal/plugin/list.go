package plugin

import (
	"fmt"
	"os"
	"os/exec"
)

func List() {

	cmd := exec.Command(
		"PluginsManagerCMD",
		"status",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		fmt.Println("Plugin Manager not installed.")
	}
}