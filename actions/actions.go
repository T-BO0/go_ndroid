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

// Swipe function to swipe on the screen.
func (node Node) Swipe(x1, y1, x2, y2 int) error {
	err := adb.Swipe(x1, y1, x2, y2)
	if err != nil {
		return fmt.Errorf("failed to swipe: %v", err)
	}

	return nil
}

// DoubleClick function to double click on an element.
func (node Node) DoubleClick() error {
	err := node.Click()
	if err != nil {
		return fmt.Errorf("failed to click: %v", err)
	}
	err = node.Click() // second click
	if err != nil {
		return fmt.Errorf("failed to click second time: %v", err)
	}
	return nil
}

// LongClick function to long click on an element with given duration.
func (node *Node) LongClick(timeout int) error {
	x, y, err := node.calculateMiddlePoint()
	if err != nil {
		return fmt.Errorf("failed to calculate middle point: %v", err)
	}

	err = adb.SwipeWithDuration(x, y, x, y, timeout)
	if err != nil {
		return fmt.Errorf("failed to long tap on element: %v", err)
	}

	return nil
}

// InsertText function to insert text into an element.
func (node *Node) InsertText(text string) error {
	err := node.Click()
	if err != nil {
		return fmt.Errorf("failed to focus on element: %v", err)
	}

	err = adb.InputText(text)
	if err != nil {
		return fmt.Errorf("failed to input text: %v", err)
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

// ElementIsVisibleBasedOnContent function to check if an element is visible based on its content Desc.
func (node *Node) ElementIsVisibleBasedOnContentDesc(content string, timeout int) bool {
	for i := 0; i < timeout; i++ {
		_, err := node.FindElementByContentDesc(content)
		if err == nil {
			return true
		}
		time.Sleep(1 * time.Second)
	}
	return false
}

// MustClick function to click on an element and panic if an error occurs.
func (node *Node) MustClick() {
	err := node.Click()
	if err != nil {
		panic(err)
	}
}

// MustSwipe function to swipe on the screen and panic if an error occurs.
func MustSwipe(x1, y1, x2, y2 int) {
	err := adb.Swipe(x1, y1, x2, y2)
	if err != nil {
		panic(err)
	}
}

// MustDoubleClick function to double click on an element and panic if an error occurs.
func (node *Node) MustDoubleClick() {
	err := node.DoubleClick()
	if err != nil {
		panic(err)
	}
}

// MustLongClick function to long click on an element with given duration and panic if an error occurs.
func (node *Node) MustLongClick(timeout int) {
	err := node.LongClick(timeout)
	if err != nil {
		panic(err)
	}
}

// MustInsertText function to insert text into an element and panic if an error occurs.
func (node *Node) MustInsertText(text string) {
	err := node.InsertText(text)
	if err != nil {
		panic(err)
	}
}
