package actions

// ElementIsCheckable function to check if an element is checkable based on its ID.
func (node *Node) ElementIsCheckable() bool {
	for _, atr := range node.Attributes {
		if atr.Name.Local == "checkable" && atr.Value == "true" {
			return true
		}
	}
	return false
}

// ElementIsChecked function to check if an element is checked.
func (node *Node) ElementIsChecked() bool {
	for _, atr := range node.Attributes {
		if atr.Name.Local == "checked" && atr.Value == "true" {
			return true
		}
	}
	return false
}

// ElementIsClickable function to check if an element is clickable.
func (node *Node) ElementIsClickable() bool {
	for _, atr := range node.Attributes {
		if atr.Name.Local == "clickable" && atr.Value == "true" {
			return true
		}
	}
	return false
}

// ElementIsEnabled function to check if an element is enabled.
func (node *Node) ElementIsEnabled() bool {
	for _, atr := range node.Attributes {
		if atr.Name.Local == "enabled" && atr.Value == "true" {
			return true
		}
	}
	return false
}

// ElementIsFocusable function to check if an element is focusable.
func (node *Node) ElementIsFocusable() bool {
	for _, atr := range node.Attributes {
		if atr.Name.Local == "focusable" && atr.Value == "true" {
			return true
		}
	}
	return false
}

// ElementIsFocused function to check if an element is focused.
func (node *Node) ElementIsFocused() bool {
	for _, atr := range node.Attributes {
		if atr.Name.Local == "focused" && atr.Value == "true" {
			return true
		}
	}
	return false
}

// ElementIsLongClickable function to check if an element is long clickable.
func (node *Node) ElementIsLongClickable() bool {
	for _, atr := range node.Attributes {
		if atr.Name.Local == "long-clickable" && atr.Value == "true" {
			return true
		}
	}
	return false
}

// ElementIsPassword function to check if an element is a password field.
func (node *Node) ElementIsPassword() bool {
	for _, atr := range node.Attributes {
		if atr.Name.Local == "password" && atr.Value == "true" {
			return true
		}
	}
	return false
}

// ElementIsScrollable function to check if an element is scrollable.
func (node *Node) ElementIsScrollable() bool {
	for _, atr := range node.Attributes {
		if atr.Name.Local == "scrollable" && atr.Value == "true" {
			return true
		}
	}
	return false
}

// ElementIsSelected function to check if an element is selected.
func (node *Node) ElementIsSelected() bool {
	for _, atr := range node.Attributes {
		if atr.Name.Local == "selected" && atr.Value == "true" {
			return true
		}
	}
	return false
}
