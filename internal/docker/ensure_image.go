package docker

import "fmt"

func EnsureImage(image string) error {

	if ImageExists(image) {
		return nil
	}

	fmt.Println("Downloading", image)

	return PullImage(image)
}