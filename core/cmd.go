package core

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// Function to run an ADB command.
// Take arguments comma separated.
// Example: RunAdbCommand(true, "input", "tap", "100", "100")
func RunAdbCommand(throughADB bool, args ...string) (string, error) {
	var cmd *exec.Cmd
	// Create the command object
	if throughADB {
		args = append([]string{"shell"}, args...)
		cmd = exec.Command("adb", args...)
		fmt.Println(cmd.String())
	} else {
		cmd = exec.Command(args[0], args[1:]...)
		fmt.Println(cmd.String())
	}
	// Capture standard output and standard error
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error: %v, stderr: %s", err, stderr.String())
	}

	// Return the output
	result := strings.TrimSpace(stdout.String())
	return result, nil
}
