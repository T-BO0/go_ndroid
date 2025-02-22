# go_ndroid

go_ndroid is a Go package for automating interactions with Android devices using ADB (Android Debug Bridge). It provides various functionalities to interact with Android devices, such as running shell commands, installing/uninstalling apps, interacting with UI elements, and capturing device logs.

## Features

- List connected devices
- Install and uninstall apps
- Launch and stop apps
- Clear app data
- Retrieve device properties
- Push and pull files
- Capture logcat logs
- Reboot device
- Retrieve battery level
- Take screenshots
- Record screen
- Tap and swipe on the screen
- Interact with UI elements

## Installation

To install the package, use the following command:

```sh
go get github.com/T-BO0/go_ndroid
```

## Usage

List available devices

```go
package main

import (
    "fmt"
    "github.com/T-BO0/go_ndroid/core"
)

func main() {
    output, err := core.RunAdbCommand("devices")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Output:", output)
    }
}
```

Work with app

```go
package main

import (
	"github.com/T-BO0/go_ndroid/actions"
	"github.com/T-BO0/go_ndroid/adb"
)

func main() {
	// Create a new Adb instance
	drv := adb.NewAdb("<device-id>", "<app-package>", "<main-activity>", false)

	// Launch the app
	drv.LaunchApp()

	// Get the root element (XML structure of the current screen)
	root, err := actions.GetPage()
	if err != nil {
		panic(err)
	}

	// Check if the element with the id "<element-id>" is visible in the root element (30 is timeout in seconds)
	visible := root.ElementIsVisibleBasedOnID("<element-id>", 30)
	if !visible {
		panic("Element not visible")
	}

	// Find the element with the id "<element-id>" and send the text "test" to it
	err = root.MustGetElementById("<element-Id>").InsertText("test")
	if err != nil {
		panic(err)
	}

	// Simulate a key event (Enter)
	err = adb.SendKeyevent(adb.KEYCODE_ENTER)
	if err != nil {
		panic(err)
	}

	// Get structure of the new page
	newRoot, err := actions.GetPage()
	if err != nil {
		panic(err)
	}

	// Check if the element with the id "<element-id>" is visible in the new root element (30 is timeout in seconds)
	visible = newRoot.ElementIsVisibleBasedOnText("<element-text>", 30)
	if !visible {
		panic("Element not visible")
	}

	// Find the element with the text "<element-text>" and click on it
	err = newRoot.MustGetElementByText("<element-text>").Click()
	if err != nil {
		panic(err)
	}

	// close the app
	drv.QuitApp()
}
```

Web

```go
package main

import (
	"github.com/T-BO0/go_ndroid/actions"
	"github.com/T-BO0/go_ndroid/adb"
)

func main() {
	// Create a new ADB instance
	drv := new(adb.Adb)

	// Open Chrome and navigate to google.com
	page, err := actions.OpenChrome("https://www.google.com", drv)
	if err != nil {
		panic(err)
	}

	// Find the search bar
	searchBar := page.MustGetElementByClass("android.widget.EditText")

	// Type "Hello, World!" into the search bar
	searchBar.MustInsertText("Hello, World!")

	// Press the enter key
	adb.SendKeyevent(adb.KEYCODE_ENTER)

	// Close chrome
	drv.QuitApp()
}

```
