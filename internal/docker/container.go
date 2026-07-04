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

func RunContainer(
	name string,
	image string,
	network string,
	volumes []string,
	ports ...string,
) error {

	if ContainerExists(name) {
		return nil
	}

	args := []string{
		"run",
		"-d",
		"--name", name,
		"--network", network,
	}

	for _, v := range volumes {
		args = append(args, "-v", v)
	}

	for _, p := range ports {
		args = append(args, "-p", p)
	}

	args = append(args, image)

	cmd := exec.Command("docker", args...)

	return cmd.Run()
}

func StartContainer(name, image string) error {
	return RunContainer(name, image, "bridge", nil)
}