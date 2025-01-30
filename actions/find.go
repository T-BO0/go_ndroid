package actions

import (
	"fmt"
	"strings"
	"time"

	"github.com/T-BO0/go_ndroid/core"
)

// FindElementByText function to find an element by text.
// The function will search for the element for duration of timeout seconds.
func FindElementByText(text string, timeout int) (string, error) {
	core.DumpXMLToFile()
	out, err := core.ReadXML()
	if err != nil {
		return "", err
	}

	//check if out contains text if not retry for 15 seconds
	for i := 0; i < timeout; i++ {
		if !strings.Contains(out, text) {
			core.DumpXMLToFile()
			out, err = core.ReadXML()
			if err != nil {
				return "", err
			}
			time.Sleep(1 * time.Second)
		} else {
			break
		}
	}
	elementArray := strings.Split(out, "<node")
	for _, doc := range elementArray {
		if strings.Contains(doc, text) {
			return doc, nil
		}
	}

	return "", fmt.Errorf("element with text %s not found", text)
}
