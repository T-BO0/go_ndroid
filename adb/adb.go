package adb

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/T-BO0/go_ndroid/core"
)

// Adb struct contains the DeviceID and PackageName which are used for device and app interactions.
type Adb struct {
	DeviceID           string // Unique ID of the Android device
	PackageName        string // The package name of the app being interacted with
	MainActivity       string // The main activity of the app being interacted with
	WithAllPermissions bool   // Grant all permissions to the app
}

// NewAdb creates a new Adb struct with the specified DeviceID, PackageName, MainActivity, and WithAllPermissions.
// NOTE withAllPermissions requires root access. make sure to have root access before using it. otherwise run withAllPermison = false.
func NewAdb(deviceID, packageName, mainActivity string, withAllPermissions bool) *Adb {
	return &Adb{
		DeviceID:           deviceID,
		PackageName:        packageName,
		MainActivity:       mainActivity,
		WithAllPermissions: withAllPermissions,
	}
}

// ListDevices lists all the connected devices and returns their device IDs.
func (adb *Adb) ListDevices() ([]string, error) {
	out, err := core.RunAdbCommand("devices")
	if err != nil {
		return nil, fmt.Errorf("failed to get list of devices - %v", err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	var devices []string
	for scanner.Scan() {
		line := scanner.Text()

		// Skip the header and empty lines
		if strings.Contains(line, "List of devices attached") || strings.TrimSpace(line) == "" {
			continue
		}

		// Split the line and check if it's a device
		parts := strings.Fields(line)
		if len(parts) > 1 && parts[1] == "device" {
			devices = append(devices, parts[0])
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to parse adb devices output: %v", err)
	}

	return devices, nil
}

// IsDeviceOnline checks whether a device with the given deviceID is currently online.
func (adb *Adb) IsDeviceOnline(deviceID string) (bool, error) {
	devices, err := adb.ListDevices()
	if err != nil {
		return false, err
	}

	// Check if the deviceID is in the list of connected devices
	for _, v := range devices {
		if v == deviceID {
			return true, nil
		}
	}

	return false, nil
}

// InstallApp installs an APK on the device.
func (adb *Adb) InstallApp(apkPath string) error {
	_, err := core.RunAdbCommand("install", apkPath)
	if err != nil {
		return fmt.Errorf("failed to install package - %v", err)
	}
	return nil
}

// UninstallApp uninstalls the app identified by PackageName.
func (adb *Adb) UninstallApp() error {
	_, err := core.RunAdbCommand("uninstall", adb.PackageName)
	if err != nil {
		return fmt.Errorf("failed to uninstall package %s - %v", adb.PackageName, err)
	}
	return nil
}

// LaunchApp launches the app with the specified activity.
// If WithAllPermissions is set to true, all permissions will be granted to the app.
func (adb *Adb) LaunchApp() error {
	if adb.WithAllPermissions {
		err := adb.GrantAllPermissions()
		if err != nil {
			return fmt.Errorf("failed to grant all permissions to app - %v", err)
		}
	}

	_, err := core.RunAdbCommand("shell", "am", "start", "-n", fmt.Sprintf("%s/%s", adb.PackageName, adb.MainActivity), "-S")
	if err != nil {
		return fmt.Errorf("failed to launch app - %v", err)
	}
	return nil
}

// GrantAllPermissions grants all permissions to the app.
// NOTE - This requires root access.
func (adb *Adb) GrantAllPermissions() error {
	err := core.GrantAllPermissions(adb.PackageName)
	if err != nil {
		return fmt.Errorf("failed to grant all permissions to app - %v", err)
	}
	return nil
}

// GrantPermission grants a specific permission to the app.
// NOTE - This requires root access.
func (adb *Adb) GrantPermission(permission AndroidPermission) error {
	err := core.GrantPermission(adb.PackageName, PermissionStrings[permission])
	if err != nil {
		return fmt.Errorf("failed to grant permission %s - %v", PermissionStrings[permission], err)
	}
	return nil
}

// RevokePermission revokes a specific permission from the app.
// NOTE - This requires root access.
func (adb *Adb) RevokePermission(permission string) error {
	err := core.RevokePermission(adb.PackageName, permission)
	if err != nil {
		return fmt.Errorf("failed to revoke permission %s - %v", permission, err)
	}
	return nil
}

// StopApp stops the app identified by PackageName.
func (adb *Adb) StopApp() error {
	_, err := core.RunAdbCommand("shell", "am", "force-stop", adb.PackageName)
	if err != nil {
		return fmt.Errorf("failed to stop app - %v", err)
	}
	return nil
}

// ClearAppData clears the app data for the app identified by PackageName.
func (adb *Adb) ClearAppData() error {
	_, err := core.RunAdbCommand("shell", "pm", "clear", adb.PackageName)
	if err != nil {
		return fmt.Errorf("failed to clear app data - %v", err)
	}
	return nil
}

// GetDeviceProperty retrieves a specific device property by its name.
func (adb *Adb) GetDeviceProperty(property string) (string, error) {
	out, err := core.RunAdbCommand("shell", "getprop", property)
	if err != nil {
		return "", fmt.Errorf("failed to read property %s: %v", property, err)
	}
	return strings.Trim(out, " "), nil
}

// PushFile pushes a local file to a remote path on the device.
func (adb *Adb) PushFile(localPath, remotePath string) error {
	_, err := core.RunAdbCommand("push", localPath, remotePath)
	if err != nil {
		return fmt.Errorf("failed to push file from %s to %s", localPath, remotePath)
	}
	return nil
}

// PullFile pulls a file from the device to a local path.
func (adb *Adb) PullFile(remotePath, localPath string) error {
	_, err := core.RunAdbCommand("pull", remotePath, localPath)
	if err != nil {
		return fmt.Errorf("failed to pull file from %s to %s", remotePath, localPath)
	}
	return nil
}

// StartLogcat starts capturing logcat logs and saves them to the specified file.
func (adb *Adb) StartLogcat(logFilePath string) error {
	_, err := core.RunAdbCommand("logcat", ">", logFilePath)
	if err != nil {
		return fmt.Errorf("failed to start logcat to file %s", logFilePath)
	}
	return nil
}

// ClearLogcat clears the logcat logs on the device.
func (adb *Adb) ClearLogcat() error {
	_, err := core.RunAdbCommand("logcat", "-c")
	if err != nil {
		return fmt.Errorf("failed to clear logcat logs - %v", err)
	}
	return nil
}

// RebootDevice reboots the device.
func (adb *Adb) RebootDevice() error {
	_, err := core.RunAdbCommand("reboot")
	if err != nil {
		return fmt.Errorf("failed to reboot device - %v", err)
	}
	return nil
}

// GetBatteryLevel retrieves the battery level of the device.
func (adb *Adb) GetBatteryLevel() (int, error) {
	out, err := core.RunAdbCommand("shell", "dumpsys", "battery", "|", "grep", "level")
	if err != nil {
		return 0, fmt.Errorf("failed to get battery level - %v", err)
	}

	// Extract digits from the output using a regular expression
	re := regexp.MustCompile(`\d+`)
	digits := re.FindString(out)

	// Convert the extracted string digits to an integer
	batteryLevel, err := strconv.Atoi(digits)
	if err != nil {
		return 0, fmt.Errorf("failed to convert battery level to integer - %v", err)
	}

	return batteryLevel, nil
}

// TakeScreenshot takes a screenshot and saves it to the specified path.
func (adb *Adb) TakeScreenshot(savePath string) error {
	_, err := core.RunAdbCommand("shell", "screencap", "-p", savePath)
	if err != nil {
		return fmt.Errorf("failed to take screenshot and save to %s: %v", savePath, err)
	}
	return nil
}

// StartScreenRecording starts recording the screen and saves it to the specified path.
func (adb *Adb) StartScreenRecording(savePath string) error {
	_, err := core.RunAdbCommand("shell", "screenrecord", savePath)
	if err != nil {
		return fmt.Errorf("failed to start screen recording on file %s: %v", savePath, err)
	}
	return nil
}

// StopScreenRecording stops the screen recording process.
func (adb *Adb) StopScreenRecording() error {
	_, err := core.RunAdbCommand("shell", "pkill", "-l", "9", "screenrecord")
	if err != nil {
		return fmt.Errorf("failed to stop screen recording - %v", err)
	}
	return nil
}

// Tap taps on the screen at the specified coordinates.
func Tap(x, y int) error {
	_, err := core.RunAdbCommand("shell", "input", "tap", fmt.Sprintf("%d", x), fmt.Sprintf("%d", y))
	if err != nil {
		return fmt.Errorf("failed to tap on screen at coordinates (%d, %d) - %v", x, y, err)
	}
	return nil
}

// Swipe swipes on the screen from the specified start to end coordinates.
func Swipe(x1, y1, x2, y2 int) error {
	_, err := core.RunAdbCommand("shell", "input", "swipe", fmt.Sprintf("%d", x1), fmt.Sprintf("%d", y1), fmt.Sprintf("%d", x2), fmt.Sprintf("%d", y2))
	if err != nil {
		return fmt.Errorf("failed to swipe on screen from (%d, %d) to (%d, %d) - %v", x1, y1, x2, y2, err)
	}
	return nil
}

// SwipeWithDuration swipes on the screen from the specified start to end coordinates with the specified duration.
func SwipeWithDuration(x1, y1, x2, y2, duration int) error {
	_, err := core.RunAdbCommand("shell", "input", "swipe", fmt.Sprintf("%d", x1), fmt.Sprintf("%d", y1), fmt.Sprintf("%d", x2), fmt.Sprintf("%d", y2), fmt.Sprintf("%d", duration))
	if err != nil {
		return fmt.Errorf("failed to swipe on screen from (%d, %d) to (%d, %d) with duration %d - %v", x1, y1, x2, y2, duration, err)
	}
	return nil
}

// InputText inputs the specified text on the device.
func InputText(text string) error {
	_, err := core.RunAdbCommand("shell", "input", "text", text)
	if err != nil {
		return fmt.Errorf("failed to input text %s - %v", text, err)
	}
	return nil
}

// SendKeyevent sends a key event to the device.
func SendKeyevent(keyEvent KeyEvent) error {
	_, err := core.RunAdbCommand("shell", "input", "keyevent", KeyEvents[keyEvent])
	if err != nil {
		return fmt.Errorf("failed to send key event %d - %v", keyEvent, err)
	}
	return nil
}
