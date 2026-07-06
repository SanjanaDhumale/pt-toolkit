package docker

import "os/exec"

func ImageExists(image string) bool {

	cmd := exec.Command(
		"docker",
		"image",
		"inspect",
		image,
	)

	return cmd.Run() == nil
}

