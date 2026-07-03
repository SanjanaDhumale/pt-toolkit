package installer

import (
	"fmt"

	"github.com/SanjanaDhumale/pt-toolkit/internal/system"
)

func CheckDocker() {

	fmt.Println("Checking Docker...")

	version, err := system.CheckDocker()

	if err != nil {
		fmt.Println("✗ Docker Not Found")
		return
	}

	fmt.Println("✓ Docker Ready")
	fmt.Println("Version:", version)
}