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

// Function to read the XML of the current screen from the file /data/local/tmp/uidump.xml
func ReadXML() (string, error) {
	out, err := RunAdbCommand("shell", "cat", "/data/local/tmp/uidump.xml")
	if err != nil {
		return "", err
	}
	return out, err
}

// Function to grant a permission to an app
func GrantPermission(packageName, permission string) error {
	_, err := RunAdbCommand("shell", "pm", "grant", packageName, permission)
	if err != nil {
		return fmt.Errorf("error granting permission: %v", err)
	}
	return nil
}

// Function to revoke a permission from an app
func RevokePermission(packageName, permission string) error {
	_, err := RunAdbCommand("shell", "pm", "revoke", packageName, permission)
	if err != nil {
		return fmt.Errorf("error revoking permission: %v", err)
	}
	return nil
}

// Function to grant all permissions to an app
func GrantAllPermissions(packageName string) error {
	// Get the list of permissions
	permissions, err := RunAdbCommand("shell", "dumpsys", "package", packageName)
	if err != nil {
		return fmt.Errorf("error getting permissions: %v", err)
	}

	// Extract permission names
	lines := strings.Split(permissions, "\n")
	for _, line := range lines {
		if strings.Contains(line, "android.permission") {
			parts := strings.Fields(line)
			if len(parts) > 1 {
				permission := parts[1]
				// Grant the permission
				err := GrantPermission(packageName, permission)
				if err != nil {
					return fmt.Errorf("error granting permission %s: %v", permission, err)
				}
			}
		}
	}
	return nil
}
