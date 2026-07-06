package report

import (
	"os"
	"path/filepath"
	"time"
)

func CreateReportFolder(workspace string) (string, error) {

	now := time.Now().Format("2006-01-02_150405")

	path := filepath.Join(
		"workspace",
		workspace,
		"reports",
		now,
	)

	err := os.MkdirAll(path, os.ModePerm)

	return path, err
}
