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
