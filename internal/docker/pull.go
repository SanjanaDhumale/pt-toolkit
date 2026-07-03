package docker

import "os/exec"

func PullImage(image string) error {

	cmd := exec.Command(
		"docker",
		"pull",
		image,
	)

	return cmd.Run()
}