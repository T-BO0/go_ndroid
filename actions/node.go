package actions

import (
	"encoding/xml"
	"fmt"

	"github.com/T-BO0/go_ndroid/core"
)

// Bounds represents the bounds of an element
type bounds struct {
	Left   int `xml:"left,attr"`
	Top    int `xml:"top,attr"`
	Right  int `xml:"right,attr"`
	Bottom int `xml:"bottom,attr"`
}

// Node represents a generic XML node
type Node struct {
	XMLName    xml.Name
	Attributes []xml.Attr `xml:",any,attr"`
	Children   []Node     `xml:",any"`
	Content    string     `xml:",chardata"`
}

// GetPage function to get the root node of the XML
func GetPage() (*Node, error) {
	if err := core.DumpXMLToFile(); err != nil {
		return &Node{}, err
	}

	xmlDoc, err := core.ReadXML()
	if err != nil {
		return &Node{}, err
	}

	xmlData := []byte(xmlDoc)

	var root Node
	err = xml.Unmarshal(xmlData, &root)
	if err != nil {
		return &Node{}, err
	}

	return &root, nil
}

// FindElementById function to find an element by id.
func (node *Node) FindElementById(resourceId string) (*Node, error) {
	for _, atrr := range node.Attributes {
		if atrr.Name.Local == "resource-id" && atrr.Value == resourceId {
			return node, nil
		}
	}

	for _, child := range node.Children {
		if found, err := child.FindElementById(resourceId); err == nil {
			return found, nil
		}
	}

	return &Node{}, fmt.Errorf("element with id %s not found", resourceId)
}

// FindElementByText function to find an element by text.
func (node *Node) FindElementByText(text string) (*Node, error) {
	for _, atrr := range node.Attributes {
		if atrr.Name.Local == "text" && atrr.Value == text {
			return node, nil
		}
	}

	for _, child := range node.Children {
		if found, err := child.FindElementByText(text); err == nil {
			return found, nil
		}
	}

	return &Node{}, fmt.Errorf("element with text %s not found", text)
}

// FindElementByContentDesc function to find an element by content-desc.
func (node *Node) FindElementByContentDesc(contentDesc string) (*Node, error) {
	for _, atrr := range node.Attributes {
		if atrr.Name.Local == "content-desc" && atrr.Value == contentDesc {
			return node, nil
		}
	}

	for _, child := range node.Children {
		if found, err := child.FindElementByContentDesc(contentDesc); err == nil {
			return found, nil
		}
	}

	return &Node{}, fmt.Errorf("element with content-desc %s not found", contentDesc)
}
