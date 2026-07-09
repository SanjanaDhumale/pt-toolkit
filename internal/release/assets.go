package release

import (
	"fmt"
	"os"
	"path/filepath"
)

func CopyReleaseAssets() error {

	fmt.Println("Step 4/5 : Copying release assets...")

	// README
	if err := CopyFile(
		"README.md",
		filepath.Join(ReleaseFolder, "README.md"),
	); err != nil {
		return err
	}

	// LICENSE (only if it exists)
	if _, err := os.Stat("LICENSE"); err == nil {
		if err := CopyFile(
			"LICENSE",
			filepath.Join(ReleaseFolder, "LICENSE"),
		); err != nil {
			return err
		}
	}

	// docs/
	if _, err := os.Stat("docs"); err == nil {
		if err := CopyDirectory(
			"docs",
			filepath.Join(ReleaseFolder, "docs"),
		); err != nil {
			return err
		}
	}

	// docker/
	if _, err := os.Stat("docker"); err == nil {
		if err := CopyDirectory(
			"docker",
			filepath.Join(ReleaseFolder, "docker"),
		); err != nil {
			return err
		}
	}

	fmt.Println("✓ Release Assets Copied")
	fmt.Println()

	return nil
}
