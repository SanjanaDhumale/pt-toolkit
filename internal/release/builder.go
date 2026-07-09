package release

import (
	"fmt"
)

func BuildRelease() error {

	fmt.Println()
	fmt.Println("PT Toolkit Release Builder")
	fmt.Println("===================================")
	fmt.Println()

	if err := BuildExecutable(); err != nil {
		return err
	}

	if err := CreateReleaseFolder(); err != nil {
		return err
	}

	fmt.Println("Step 3/5 : Copying executable...")

	if err := CopyFile(
		"build/ptctl.exe",
		ReleaseFolder+"/ptctl.exe",
	); err != nil {
		return err
	}

	fmt.Println("✓ Executable Copied")
	fmt.Println()

	if err := CopyReleaseAssets(); err != nil {
		return err
	}

	if err := PackageRelease(); err != nil {
		return err
	}

	return nil
}
