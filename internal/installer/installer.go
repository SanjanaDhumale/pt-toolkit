package installer

import (
	"fmt"

	"github.com/SanjanaDhumale/pt-toolkit/internal/config"
)

func Install() {

	CheckDocker()

	InstallImages()

	CreateFolders()

	cfg := config.Default()

	if err := config.Save(cfg); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()
	fmt.Println("✓ Installation Completed")
}