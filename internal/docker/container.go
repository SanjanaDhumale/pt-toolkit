package docker

import (
	"os/exec"
)

func ContainerExists(name string) bool {

	cmd := exec.Command(
		"docker",
		"container",
		"inspect",
		name,
	)

	return cmd.Run() == nil
}

func RunContainer(name, image, network string, ports ...string) error {

	if ContainerExists(name) {
		return nil
	}

	args := []string{
		"run",
		"-d",
		"--name", name,
		"--network", network,
	}

	for _, p := range ports {
		args = append(args, "-p", p)
	}

	args = append(args, image)

	cmd := exec.Command("docker", args...)

	return cmd.Run()
}

// --------------------------------------------------------------------
// Backward Compatibility
// Old service commands still call StartContainer().
// This wrapper lets them continue working until we migrate completely.
// --------------------------------------------------------------------

func StartContainer(name, image string) error {
	return RunContainer(name, image, "bridge")
}