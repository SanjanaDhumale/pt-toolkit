package execution

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/SanjanaDhumale/pt-toolkit/internal/workspace"
)

func RunJMeter(script string) error {

	ws, err := workspace.CurrentWorkspace()

	if err != nil {
		return err
	}

	scriptPath := filepath.Join(
		"workspace",
		ws,
		"scripts",
		"jmeter",
		script,
	)

	if _, err := os.Stat(scriptPath); err != nil {
		return fmt.Errorf("script not found: %s", scriptPath)
	}

	fmt.Println()
	fmt.Println("====================================")
	fmt.Println(" Running JMeter")
	fmt.Println("====================================")
	fmt.Println()

	fmt.Println("Workspace :", ws)
	fmt.Println("Script    :", script)
	fmt.Println("Location  :", scriptPath)

	fmt.Println()
	fmt.Println("✓ Script Found")

	return nil
}
