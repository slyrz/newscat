package html

import (
	"code.google.com/p/go.net/html"
	"errors"
	"github.com/slyrz/newscat/util"
	"strings"
)

// Errors returned by the NewChunk function.
var (
	ErrNoParent = errors.New("no parent")
	ErrNoText   = errors.New("no text")
	ErrNoBlock  = errors.New("no block node after parent")
)

// A Chunk is a chunk of consecutive text found in the HTML document.
// It combines the content of one or more html.TextNodes. Whitespace is
// ignored, but inter-word separation is preserved. Therefore each Chunk
// must contain actual text and whitespace-only html.TextNodes don't
// result in Chunks.
type Chunk struct {
	Prev      *Chunk     // previous chunk
	Next      *Chunk     // next chunk
	Text      *util.Text // text of this chunk
	Base      *html.Node // element node which contained this chunk
	Block     *html.Node // parent block node of base node
	Container *html.Node // parent block node of block node
	Classes   []string   // list of classes this chunk belongs to
	Ancestors int        // bitmask of the ancestors of this chunk
	LinkText  float32    // link text to normal text ratio.
}

func getParentBlock(n *html.Node) *html.Node {
	// Keep ascending as long as the node points to an HTML inline element.
	// We stop at the first block-level element. The list of inline elements
	// was taken from:
	//   https://developer.mozilla.org/en-US/docs/HTML/Inline_elements
	for ; n != nil && n.Parent != nil; n = n.Parent {
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
	// If a TextNode was passed, use the parent ElementNode for the
	// base field.
	case html.TextNode:
		// We don't allow orphaned Chunks.
		if n.Parent == nil {
			return nil, ErrNoParent
		}
		chunk.Base = n.Parent
	}

	// Write the text of all TextNodes of n to chunk.Text.
	iterateText(n, chunk.Text.WriteString)

	// Don't produce Chunks without text.
	if chunk.Text.Len() == 0 {
		return nil, ErrNoText
	}

	// Now we detect the HTML block and container of the base node. The block
	// is the first block-level element found when ascending from base node.
	// The container is the first block-level element found when ascending
	// from the block's parent.
	//
	// Example
	//
	// Base node is a block-level element:
	//
	//   <div>                        <- Container
	//     <p>Hello World</p>         <- Base & Block
	//   </div>
	//
	// Base node is not a block-level element:
	//
	//   <div>                         <- Container
	//     <p>                         <- Block
	//       <span>
	//         <i>Hello World</i>      <- Base
	//       </span>
	//     </p>
	//   </div>
	if block := getParentBlock(chunk.Base); block != nil {
		chunk.Block = block
	} else {
		return nil, ErrNoBlock
	}

	// If there happens to be no block-level element after the block's parent,
	// use block as container as well. This ensures that the container field
	// is never nil and we avoid nil pointer handling in our code.
	if container := getParentBlock(chunk.Block.Parent); container != nil {
		chunk.Container = container
	} else {
		chunk.Container = chunk.Block
	}

	// Remember the ancestors in our chunk.
	chunk.Ancestors = doc.ancestors

	// Calculate the ratio between text inside links and text outside links
	// for the current element's block node. This is useful to determine the
	// quality of a link. Links used as cross references inside the doc
	// content have a small link text to text ratio,
	//
	//   <p>Long text .... <a>short text</a> ... </p>
	//
	// whereas related content / navigation links have a high link text
	// to text ratio:
	//
	//   <li><a>See also: ...</a></li>
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
	// the content class. Most IDs aren't really meaningful, so no IDs here.
	chunk.Classes = make([]string, 0)

	// Ascend parent nodes until we found a class attribute and some
	// microdata.
	haveClass := false
	haveMicro := false
	for prev := chunk.Base; prev != nil; prev = prev.Parent {
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
			// The default: continue case keeps us from reaching this for
			// attributes we are not interested in.
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

// Returns a list of strings containing the HTML element types
// of the Chunk's siblings.
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

// Returns a list of strings containing the HTML element types
// of the Chunk's children.
func (ch *Chunk) GetChildTypes() []string {
	result := make([]string, 0, 8)
	for s := ch.Base.FirstChild; s != nil; s = s.NextSibling {
		if s.Type == html.ElementNode {
			result = append(result, s.Data)
		}
	}
	return result
}
