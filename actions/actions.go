package actions

import (
	"fmt"
	"time"

	"github.com/T-BO0/go_ndroid/adb"
)

// Click function to click on an element.
func (node *Node) Click() error {
	x, y, err := node.calculateMiddlePoint()
	if err != nil {
		return fmt.Errorf("failed to calculate middle point: %v", err)
	}

	err = adb.Tap(x, y)
	if err != nil {
		return fmt.Errorf("failed to tap on element: %v", err)
	}

	return nil
}

func (node Node) Swipe(x1, y1, x2, y2 int) error {
	err := adb.Swipe(x1, y1, x2, y2)
	if err != nil {
		return fmt.Errorf("failed to swipe: %v", err)
	}

	return nil
}

// ElementIsVisibleBasedOnID function to check if an element is visible based on its ID.
func (node *Node) ElementIsVisibleBasedOnID(resourceId string, timeout int) bool {
	for i := 0; i < timeout; i++ {
		_, err := node.FindElementById(resourceId)
		if err == nil {
			return true
		}
		time.Sleep(1 * time.Second)
	}

	return false
}

// ElementIsVisibleBasedOnText function to check if an element is visible based on its text.
func (node *Node) ElementIsVisibleBasedOnText(text string, timeout int) bool {
	for i := 0; i < timeout; i++ {
		_, err := node.FindElementByText(text)
		if err == nil {
			return true
		}
		time.Sleep(1 * time.Second)
	}

	return false
}
