package actions

import "fmt"

// Text function to get the text of an element.
func (node *Node) Text() string {
	return node.Content
}

// ResourceID function to get the resource ID of an element.
func (node *Node) ResourceID() string {
	for _, attr := range node.Attributes {
		if attr.Name.Local == "resource-id" {
			return attr.Value
		}
	}
	return ""
}

// Bounds function to get the bounds of an element.
func (node *Node) Bounds() (left, top, right, bottom int) {
	for _, attr := range node.Attributes {
		if attr.Name.Local == "bounds" {
			fmt.Sscanf(attr.Value, "[%d,%d][%d,%d]", &left, &top, &right, &bottom)
		}
	}
	return left, top, right, bottom
}

// IsClickable function to check if an element is clickable.
func (node *Node) IsClickable() bool {
	for _, attr := range node.Attributes {
		if attr.Name.Local == "clickable" {
			return attr.Value == "true"
		}
	}
	return false
}

// IsCheckable function to check if an element is checkable.
func (node *Node) IsCheckable() bool {
	for _, attr := range node.Attributes {
		if attr.Name.Local == "checkable" {
			return attr.Value == "true"
		}
	}
	return false
}

// IsChecked function to check if an element is checked.
func (node *Node) IsChecked() bool {
	for _, attr := range node.Attributes {
		if attr.Name.Local == "checked" {
			return attr.Value == "true"
		}
	}
	return false
}

// IsEnabled function to check if an element is enabled.
func (node *Node) IsEnabled() bool {
	for _, attr := range node.Attributes {
		if attr.Name.Local == "enabled" {
			return attr.Value == "true"
		}
	}
	return false
}

// IsFocusable function to check if an element is focusable.
func (node *Node) IsFocusable() bool {
	for _, attr := range node.Attributes {
		if attr.Name.Local == "focusable" {
			return attr.Value == "true"
		}
	}
	return false
}

// IsFocused function to check if an element is focused.
func (node *Node) IsFocused() bool {
	for _, attr := range node.Attributes {
		if attr.Name.Local == "focused" {
			return attr.Value == "true"
		}
	}
	return false
}

// IsScrollable function to check if an element is scrollable.
func (node *Node) IsScrollable() bool {
	for _, attr := range node.Attributes {
		if attr.Name.Local == "scrollable" {
			return attr.Value == "true"
		}
	}
	return false
}

// IsLongClickable function to check if an element is long clickable.
func (node *Node) IsLongClickable() bool {
	for _, attr := range node.Attributes {
		if attr.Name.Local == "long-clickable" {
			return attr.Value == "true"
		}
	}
	return false
}

// IsPassword function to check if an element is a password field.
func (node *Node) IsPassword() bool {
	for _, attr := range node.Attributes {
		if attr.Name.Local == "password" {
			return attr.Value == "true"
		}
	}
	return false
}

// IsSelected function to check if an element is selected.
func (node *Node) IsSelected() bool {
	for _, attr := range node.Attributes {
		if attr.Name.Local == "selected" {
			return attr.Value == "true"
		}
	}
	return false
}

// Package function to get the package name of an element.
func (node *Node) Package() string {
	for _, attr := range node.Attributes {
		if attr.Name.Local == "package" {
			return attr.Value
		}
	}
	return ""
}

// Class function to get the class name of an element.
func (node *Node) Class() string {
	for _, attr := range node.Attributes {
		if attr.Name.Local == "class" {
			return attr.Value
		}
	}
	return ""
}

// Index function to get the index of an element.
func (node *Node) Index() string {
	for _, attr := range node.Attributes {
		if attr.Name.Local == "index" {
			return attr.Value
		}
	}
	return "0"
}
