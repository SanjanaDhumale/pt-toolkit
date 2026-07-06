package execution

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/SanjanaDhumale/pt-toolkit/internal/browser"
	"github.com/SanjanaDhumale/pt-toolkit/internal/constants"
	dockerengine "github.com/SanjanaDhumale/pt-toolkit/internal/docker"
	"github.com/SanjanaDhumale/pt-toolkit/internal/report"
	"github.com/SanjanaDhumale/pt-toolkit/internal/workspace"
)

func RunJMeter(script string) error {

	// Get Current Workspace
	ws, err := workspace.CurrentWorkspace()
	if err != nil {
		return err
	}

	// Local Script Path
	scriptPath := filepath.Join(
		"workspace",
		ws,
		"scripts",
		"jmeter",
		script,
	)

	// Verify Script Exists
	if _, err := os.Stat(scriptPath); err != nil {
		return fmt.Errorf("script not found: %s", scriptPath)
	}

	// Check JMeter Container
	if dockerengine.IsRunning(constants.JMeterContainer) {
		return fmt.Errorf(
			"JMeter container is not running.\nRun:\ngo run . service start jmeter",
		)
	}

	// Container Paths
	containerScript := "/tmp/" + script
	containerResult := "/tmp/result.jtl"
	containerDashboard := "/tmp/dashboard"

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
	fmt.Println("✓ JMeter Container Running")
	fmt.Println()

	// Copy Script
	fmt.Println("Copying Script To Container...")

	err = dockerengine.CopyToContainer(
		"pt-jmeter",
		scriptPath,
		containerScript,
	)

	if err != nil {
		return err
	}

	fmt.Println("✓ Script Copied")
	fmt.Println()

	// Create Report Folder
	reportPath, err := report.CreateReportFolder(ws)
	if err != nil {
		return err
	}

	fmt.Println("Report Folder :", reportPath)
	fmt.Println()

	// Execute JMeter
	fmt.Println("Preparing Docker Execution...")

	err = dockerengine.RunJMeter(
		"pt-jmeter",
		containerScript,
		containerResult,
	)

	if err != nil {
		return err
	}

	fmt.Println("✓ Test Execution Completed")
	fmt.Println()

	// Generate Dashboard
	fmt.Println("Generating HTML Report...")

	err = dockerengine.GenerateHTMLReport("pt-jmeter")

	if err != nil {
		return err
	}

	fmt.Println("✓ HTML Report Generated")
	fmt.Println()

	// Show Dashboard Contents
	fmt.Println("Dashboard Contents")
	fmt.Println()

	err = dockerengine.Exec(
		"pt-jmeter",
		"ls",
		"-lh",
		containerDashboard,
	)

	if err != nil {
		return err
	}

	fmt.Println()

	// Copy Dashboard to Workspace
	fmt.Println("Saving Dashboard To Workspace...")

	err = dockerengine.CopyFromContainer(
		"pt-jmeter",
		containerDashboard,
		reportPath,
	)

	if err != nil {
		return err
	}

	fmt.Println("✓ Dashboard Saved")
	fmt.Println()
	fmt.Println("Opening Report...")

	indexFile := filepath.Join(
		reportPath,
		"dashboard",
		"index.html",
	)

	err = browser.Open(indexFile)

	if err != nil {
		fmt.Println("Warning: unable to open browser.")
		fmt.Println("Open manually:", indexFile)
	} else {
		fmt.Println("✓ Report Opened")
	}

	fmt.Println()
	fmt.Println("Cleaning Temporary Files...")

	err = dockerengine.Cleanup("pt-jmeter")

	if err != nil {
		return err
	}

	fmt.Println("✓ Cleanup Completed")

	fmt.Println()
	fmt.Println("====================================")
	fmt.Println(" JMeter Test Completed")
	fmt.Println("====================================")
	fmt.Println()

	return nil
}
