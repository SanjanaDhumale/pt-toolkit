package release

import (
	"fmt"
	"os"
)

const ReleaseFolder = "releases/v1.0"

func CreateReleaseFolder() error {

	fmt.Println("Step 2/5 : Creating release folder...")

	if err := os.MkdirAll(ReleaseFolder, os.ModePerm); err != nil {
		return err
	}

	fmt.Println("✓ Release Folder Ready")
	fmt.Println()

	return nil
}
