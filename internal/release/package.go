package release

import "fmt"

func PackageRelease() error {

	fmt.Println("Step 5/5 : Packaging Release...")
	fmt.Println()

	if err := BuildInstaller(); err != nil {
		return err
	}

	fmt.Println("✓ Release Package Ready")

	return nil
}
