package docker

import "os/exec"

func VolumeExists(name string) bool {

	cmd := exec.Command(
		"docker",
		"volume",
		"inspect",
		name,
	)

	return cmd.Run() == nil
}

func CreateVolume(name string) error {

	if VolumeExists(name) {
		return nil
	}

	cmd := exec.Command(
		"docker",
		"volume",
		"create",
		name,
	)

	return cmd.Run()
}