package installer

import(
	"fmt"
	"github.com/SanjanaDhumale/pt-toolkit/internal/config"
)


func Install() {

	CheckDocker()

	InstallImages()

	CreateFolders()

	cfg := config.Default()

	err := config.Save(cfg)

	if err != nil {
		return
	}

	CreateConfig()

	fmt.Println()

	fmt.Println("✓ Installation Completed")
}