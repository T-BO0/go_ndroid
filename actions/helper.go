package actions

import "fmt"

// GetBound function to get the bounds of an element.
func (node *Node) getBound() (string, error) {
	for _, atr := range node.Attributes {
		if atr.Name.Local == "bounds" {
			return atr.Value, nil
		}
	}
	return "", fmt.Errorf("bounds not found")
}

// CalculateMiddlePoint function to calculate the middle point of an element.
func (node *Node) calculateMiddlePoint() (int, int, error) {
	bounds, err := node.getBound()
	if err != nil {
		return 0, 0, err
	}

	var x1, y1, x2, y2 int
	fmt.Sscanf(bounds, "[%d,%d][%d,%d]", &x1, &y1, &x2, &y2)
	return (x1 + x2) / 2, (y1 + y2) / 2, nil
}
