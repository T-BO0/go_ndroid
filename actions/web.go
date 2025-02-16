package actions

import (
	"fmt"

	"github.com/T-BO0/go_ndroid/adb"
)

// OpenChrome opens the Chrome browser on the device.
func OpenChrome(url string, dest *adb.Adb) (*Node, error) {
	err := adb.OpenChrome(url)
	if err != nil {
		return &Node{}, fmt.Errorf("failed to open Chrome browser - %v", err)
	}

	node, err := GetPage()
	if err != nil {
		return &Node{}, fmt.Errorf("failed to get page - %v", err)
	}

	dest.DeviceID = adb.MustGetPropertie("ro.product.model")
	dest.PackageName = "com.android.chrome"
	dest.MainActivity = "com.google.android.apps.chrome.Main"
	dest.WithAllPermissions = false

	return node, nil
}
