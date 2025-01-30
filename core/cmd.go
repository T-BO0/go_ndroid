package core

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// Function to run ADB commands
func RunAdbCommand(args ...string) (string, error) {
	// Create the command object
	cmd := exec.Command("adb", args...)
	fmt.Println(cmd.String())

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

// Function to run general shell commands
func RunShellCommand(command string, args ...string) (string, error) {
	// Create the command
	cmd := exec.Command(command, args...)

	// Capture the output
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error executing command: %s, stderr: %s", err, stderr.String())
	}

	return out.String(), nil
}

// Function to dump the XML of the current screen to a file /data/local/tmp/uidump.xml
func DumpXMLToFile() error {
	_, err := RunAdbCommand("shell", "uiautomator", "dump", "/data/local/tmp/uidump.xml")
	if err != nil {
		return fmt.Errorf("error dumping XML: %v", err)
	}
	return nil
}

func ReadXML() (string, error) {
	out, err := RunAdbCommand("shell", "cat", "/data/local/tmp/uidump.xml")
	if err != nil {
		return "", err
	}
	return out, err
}
