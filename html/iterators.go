package html

import (
	"golang.org/x/net/html"
)

const (
	IterNext = iota // Keep going.
	IterSkip        // Skip the current subtree, proceed with the next sibling.
	IterStop        // Skip everything.
)

// iterateText iterates through the values of all text nodes.
func iterateText(n *html.Node, callback func(s string)) {
	if n.Type == html.TextNode {
		callback(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		iterateText(c, callback)
	}
}

// iterateNode iterates through all nodes. It further allows skipping and/or
// stopping iteration through the return value of callback.
func iterateNode(n *html.Node, callback func(s *html.Node) int) int {
	switch callback(n) {
	case IterSkip:
		return IterNext
	case IterStop:
		return IterStop
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if iterateNode(c, callback) == IterStop {
			return IterStop
		}
	}
	return IterNext
}
