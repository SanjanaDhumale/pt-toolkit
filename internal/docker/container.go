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

		if ContainerRunning(name) {
			return nil
		}

		return StartExistingContainer(name)
	}

	args := []string{
		"run",
		"-d",
		"--name", name,
		"--network", network,
	}

	for _, volume := range volumes {
		args = append(args, "-v", volume)
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

func ContainerRunning(name string) bool {

	cmd := exec.Command(
		"docker",
		"inspect",
		"-f",
		"{{.State.Running}}",
		name,
	)

	out, err := cmd.Output()

	if err != nil {
		return false
	}

	return string(out) == "true\n"
}
