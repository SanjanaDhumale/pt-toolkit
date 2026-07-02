package system

import (
	"os/exec"
	"strings"
)

// RunCommand executes a command on the operating system
// and returns its output.
func RunCommand(command string, args ...string) (string, error) {

	cmd := exec.Command(command, args...)

	output, err := cmd.CombinedOutput()

	return strings.TrimSpace(string(output)), err
}

// CheckGo checks if Go is installed.
func CheckGo() (string, error) {
	return RunCommand("go", "version")
}

// CheckDocker checks if Docker is installed.
func CheckDocker() (string, error) {
	return RunCommand("docker", "--version")
}

// CheckDockerCompose checks Docker Compose.
func CheckDockerCompose() (string, error) {
	return RunCommand("docker", "compose", "version")
}

// CheckJava checks Java installation.
func CheckJava() (string, error) {
	return RunCommand("java", "--version")
}

// CheckPython checks Python installation.
func CheckPython() (string, error) {
	return RunCommand("py", "--version")
}

// CheckGit checks Git installation.
func CheckGit() (string, error) {
	return RunCommand("git", "--version")
}

func IsDockerInstalled() bool {
	_, err := CheckDocker()
	return err == nil
}