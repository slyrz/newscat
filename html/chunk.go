package html

import (
	"code.google.com/p/go.net/html"
	"errors"
	"github.com/slyrz/newscat/util"
	"strings"
)

// A Chunk combines the text of one or more html.TextNodes.
// Whitespace-only text gets ignored and won't result in a Chunk.
type Chunk struct {
	Prev      *Chunk     // previous chunk
	Next      *Chunk     // next chunk
	Text      *util.Text // text of this chunk
	Base      *html.Node // element node which contained this chunk
	Block     *html.Node // parent block node of base node
	Container *html.Node // parent block node of block node
	Classes   []string   // list of classes this chunk belongs to
	Level     int        // depth of the element node that cointains this chunk
	Elems     int        // number of elements traversed until we reached the element node
	Ancestors int        // bitmask of the ancestors of this chunk
	LinkText  float32    // link text to normal text ratio.
}

func getParentBlock(n *html.Node) *html.Node {
	for ; n != nil && n.Parent != nil; n = n.Parent {
		// Keep ascending as long as the block node points to an HTML inline
		// element, so we stop at the first block-level element.
		// The list of inline elements was taken from
		// https://developer.mozilla.org/en-US/docs/HTML/Inline_elements
		switch n.Data {
		case "a", "abbr", "acronym", "b", "bdo", "big", "br", "button", "cite",
			"code", "dfn", "em", "i", "img", "input", "kbd", "label", "map",
			"object", "q", "samp", "script", "select", "small", "span",
			"strong", "sub", "sup", "textarea", "tt", "var":
			continue
		default:
			return n
		}
	}
	return n
}

func NewChunk(doc *Document, n *html.Node) (*Chunk, error) {
	chunk := new(Chunk)
	chunk.Text = util.NewText()

	switch n.Type {
	// If an ElementNode was passed, create Text property using all
	// TextNode children.
	case html.ElementNode:
		chunk.Base = n
		chunk.addText(n)
	// If a TextNode was passed, use the parent ElementNode for the
	// base field.
	case html.TextNode:
		// We don't allow baseless Chunks.
		if n.Parent == nil {
			return nil, errors.New("orphaned TextNode")
		}
		chunk.Base = n.Parent
		chunk.addText(n)
	}

	// We want to do text extraction, not whitespace extraction.
	if chunk.Text.Len() == 0 {
		return nil, errors.New("no text")
	}

	// Find the block level container of the base node.
	chunk.Block = getParentBlock(chunk.Base)

	// Container is the block level parent of block.
	if container := getParentBlock(chunk.Block.Parent); container != nil {
		chunk.Container = container
	} else {
		chunk.Container = chunk.Block
	}

	// Copy the document's level and element counter into chunk.
	chunk.Level = doc.level
	chunk.Elems = doc.elems
	chunk.Ancestors = doc.ancestors

	// Calculate the ratio between text inside links and text outside links
	// for the current element's block node. This is useful to determine the
	// quality of a link. Links used as cross references inside the article
	// content have a small link text to text ratio,
	//
	//	<p>Long text .... <a>short text</a> ... </p>
	//
	// whereas related content / navigation links have a high link text
	// to text ratio:
	//
	// 	<li><a>See also: ...</a></li>
	//
	linkText := doc.linkText[chunk.Block]
	normText := doc.normText[chunk.Block]
	if normText == 0 && linkText == 0 {
		chunk.LinkText = 0.0
	} else {
		chunk.LinkText = float32(linkText) / float32(linkText+normText)
	}

	// Detect the classes of the current node. We use the good old class
	// attribute and the new HTML5 microdata (itemprop attribute) to determine
	// the content class.
	chunk.Classes = make([]string, 0)

	// Ascend parent nodes until we found a class attribute and some
	// microdata.
	haveClass := false
	haveMicro := false
	for prev := chunk.Base; prev != nil; prev = prev.Parent {
		// TODO: Unlikely to happen, isn't it?
		if prev.Type != html.ElementNode {
			continue
		}
		for _, attr := range prev.Attr {
			switch {
			case !haveClass && attr.Key == "class":
				haveClass = true
			case !haveMicro && attr.Key == "itemprop":
				haveMicro = true
			default:
				continue
			}
			// The default: continue case keeps us from getting here for values
			// we are not interested in.
			for _, val := range strings.Fields(attr.Val) {
				chunk.Classes = append(chunk.Classes, val)
			}
		}
		if haveClass && haveMicro {
			break
		}
	}
	return chunk, nil
}

// Add all text from a html.Node to our chunk.
func (ch *Chunk) addText(n *html.Node) {
	switch n.Type {
	case html.TextNode:
		ch.Text.WriteString(n.Data)
	case html.ElementNode:
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			ch.addText(c)
		}
	}
}

// Return the types of the base node's siblings.
func (ch *Chunk) GetSiblingTypes() []string {
	result := make([]string, 0, 8)
	for s := ch.Base.PrevSibling; s != nil; s = s.PrevSibling {
		if s.Type == html.ElementNode {
			result = append(result, s.Data)
		}
	}
	for s := ch.Base.NextSibling; s != nil; s = s.NextSibling {
		if s.Type == html.ElementNode {
			result = append(result, s.Data)
		}
	}
	return result
}

// Return the types of the base node's children.
func (ch *Chunk) GetChildTypes() []string {
	result := make([]string, 0, 8)
	for s := ch.Base.FirstChild; s != nil; s = s.NextSibling {
		if s.Type == html.ElementNode {
			result = append(result, s.Data)
		}
	}
	return result
}

// Returns true if the base node of this chunk is a HTML heading element.
func (ch *Chunk) IsHeading() bool {
	switch ch.Base.Data {
	case "h1", "h2", "h3", "h4", "h5", "h6":
		return true
	default:
		return false
	}
}
