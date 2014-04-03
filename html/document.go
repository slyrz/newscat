package html

import (
	"code.google.com/p/go.net/html"
	"errors"
	"io"
	"regexp"
	"unicode"
)

const (
	// We remember a few special node types when descending into their
	// children.
	AncestorArticle = 1 << iota
	AncestorAside
	AncestorBlockquote
	AncestorList
)

var (
	ignorePattern = regexp.MustCompile("(?i)comment|caption|header|foot|blq-dotcom|story-feature")
)

type Document struct {
	Title  *Chunk   // the <title>...</title> text
	Chunks []*Chunk // list of all chunks found in this document

	// Unexported fields.
	root *html.Node // the <html>...</html> part
	head *html.Node // the <head>...</head> part
	body *html.Node // the <body>...</body> part

	// State variables used when collectiong chunks.
	level     int // depth of the current node
	elems     int // total number of elements visited
	ancestors int // bitmask which stores ancestor of the current node

	// Number of non-space characters inside link tags / normal tags
	// per html.ElementNode.
	linkText map[*html.Node]int // length of text inside <a></a> tags
	normText map[*html.Node]int // length of text outside <a></a> tags
}

func NewDocument(r io.Reader) (*Document, error) {
	doc := new(Document)
	doc.Chunks = make([]*Chunk, 0, 512)
	if err := doc.Parse(r); err != nil {
		return nil, err
	}
	return doc, nil
}

func (doc *Document) Parse(r io.Reader) error {
	root, err := html.Parse(r)
	if err != nil {
		return err
	}
	// Assign the fields root, head and body from the HTML page.
	doc.setNodes(root)

	// Check if <html>, <head> and <body> nodes were found.
	if doc.root == nil || doc.head == nil || doc.body == nil {
		return errors.New("Document missing <html>, <head> or <body>.")
	}

	doc.parseHead(doc.head)

	// No title found? The title plays a crucial role in detecting
	// the article content. We need it and every page should contain it.
	// So skip parsing in case we could not find a title.
	if doc.Title == nil {
		return errors.New("Document missing <title>.")
	}

	doc.level = 0
	doc.elems = 0
	doc.linkText = make(map[*html.Node]int)
	doc.normText = make(map[*html.Node]int)

	doc.countText(doc.body, false)
	doc.parseBody(doc.body)

	// Now link the chunks.
	for i := range doc.Chunks {
		if i > 0 {
			doc.Chunks[i].Prev = doc.Chunks[i-1]
		}
		if i < len(doc.Chunks)-1 {
			doc.Chunks[i].Next = doc.Chunks[i+1]
		}
	}
	return nil
}

// Assign the struct fields root, head and body from the HTML tree of node n.
//  doc.root -> <html>
//  doc.head -> <head>
//  doc.body -> <body>
func (doc *Document) setNodes(n *html.Node) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		switch c.Data {
		case "html":
			doc.root = c
			doc.setNodes(c)
		case "body":
			doc.body = c
		case "head":
			doc.head = c
		}
	}
}

// Parse the <head>...</head> part of the HTML page. This function only sets
// the title string right now.
func (doc *Document) parseHead(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "title" {
		if chunk, err := NewChunk(doc, n); err == nil {
			doc.Title = chunk
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		doc.parseHead(c)
	}
}

// countText counts the link text and the normal text per html.Node.
// "Link text" is text inside <a> tags and "normal text" is text inside
// anything but <a> tags. Of course, counting is done cumulative, so the
// numbers of a parent node include the numbers of it's child nodes.
func (doc *Document) countText(n *html.Node, insideLink bool) (linkText int, normText int) {
	linkText = 0
	normText = 0
	if n.Type == html.ElementNode && n.Data == "a" {
		insideLink = true
	}
	for s := n.FirstChild; s != nil; s = s.NextSibling {
		linkTextChild, normTextChild := doc.countText(s, insideLink)
		linkText += linkTextChild
		normText += normTextChild
	}
	if n.Type == html.TextNode {
		count := 0
		for _, rune := range n.Data {
			if unicode.IsLetter(rune) {
				count += 1
			}
		}
		if insideLink {
			linkText += count
		} else {
			normText += count
		}
	}
	doc.linkText[n] = linkText
	doc.normText[n] = normText
	return
}

// Parse the <body>...</body> part of the HTML page.
func (doc *Document) parseBody(n *html.Node) {
	doc.elems += 1
	switch n.Type {
	case html.ElementNode:
		doc.level += 1
		doc.parseBodyElement(n)
		doc.level -= 1
	case html.TextNode:
		doc.parseBodyText(n)
	}
}

func (doc *Document) parseBodyElement(n *html.Node) {
	for _, attr := range n.Attr {
		switch attr.Key {
		case "id", "class", "itemprop":
			if ignorePattern.FindStringIndex(attr.Val) != nil {
				return
			}
		}
	}

	ancestorMask := 0
	switch n.Data {
	// We convert headings to text immediately. This is easier and feasible
	// because headings don't contain many children.
	// Descending into these children and handling every TextNode separately
	// would make things unnecessary complicated and the result noisy.
	case "h1", "h2", "h3", "h4", "h5", "h6", "a":
		if chunk, err := NewChunk(doc, n); err == nil {
			doc.Chunks = append(doc.Chunks, chunk)
		}
		return
	// Elements save to ignore.
	case "address", "audio", "button", "canvas", "caption", "fieldset",
		"figcaption", "figure", "footer", "form", "frame", "header", "iframe",
		"img", "map", "menu", "nav", "noscript", "object", "option", "output",
		"script", "select", "style", "svg", "textarea", "video":
		return
	// High-level tables might be used to layout the document, so we better not
	// ignore them.
	case "table":
		if doc.level > 5 {
			return
		}
	// Now mask the element type, but only if it isn't already set.
	// If we mask a bit which was already set by one of our callers, we'd also
	// clear it at the of this function, though it actually should be cleared
	// by the caller.
	case "article":
		ancestorMask = AncestorArticle &^ doc.ancestors
	case "aside":
		ancestorMask = AncestorAside &^ doc.ancestors
	case "blockquote":
		ancestorMask = AncestorBlockquote &^ doc.ancestors
	case "ul", "ol":
		ancestorMask = AncestorList &^ doc.ancestors
	}

	doc.ancestors |= ancestorMask
	// If we reach this, we should descend.
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		doc.parseBody(c)
	}
	doc.ancestors &^= ancestorMask
}

func (doc *Document) parseBodyText(n *html.Node) {
	// Only every 6th text node contains text in our training set.
	// Since we don't need nodes that contain whitespace only, it's quite
	// beneficial to ignore these nodes.
	isWhitespace := true
	for _, rune := range n.Data {
		// The space is the greatest non printable ASCII character.
		if rune > ' ' {
			isWhitespace = false
			break
		}
	}
	if isWhitespace {
		return
	}
	if chunk, err := NewChunk(doc, n); err == nil {
		doc.Chunks = append(doc.Chunks, chunk)
	}
}

type TextStat struct {
	Words     int
	Sentences int
	Count     int
}

// GetClassStats groups the document chunks by their classes (defined by the
// class attribute of HTML nodes) and calculates TextStats for each class.
func (doc *Document) GetClassStats() map[string]*TextStat {
	result := make(map[string]*TextStat)
	for _, chunk := range doc.Chunks {
		for _, class := range chunk.Classes {
			if stat, ok := result[class]; ok {
				stat.Words += chunk.Text.Words
				stat.Sentences += chunk.Text.Sentences
				stat.Count += 1
			} else {
				result[class] = &TextStat{chunk.Text.Words, chunk.Text.Sentences, 1}
			}
		}
	}
	return result
}

// GetClusterStats groups the document chunks by common ancestors and
// calculates TextStats for each group of chunks.
func (doc *Document) GetClusterStats() map[*Chunk]*TextStat {
	// Don't ascend further than this constant.
	const maxAncestors = 3

	// Count TextStats for Chunk ancestors.
	ancestorStat := make(map[*html.Node]*TextStat)
	for _, chunk := range doc.Chunks {
		node, count := chunk.Block, 0
		for node != nil && count < maxAncestors {
			if stat, ok := ancestorStat[node]; ok {
				stat.Words += chunk.Text.Words
				stat.Sentences += chunk.Text.Sentences
				stat.Count += 1
			} else {
				ancestorStat[node] = &TextStat{chunk.Text.Words, chunk.Text.Sentences, 1}
			}
			node, count = node.Parent, count+1
		}
	}

	// Generate result.
	result := make(map[*Chunk]*TextStat)
	for _, chunk := range doc.Chunks {
		node := chunk.Block
		if node == nil {
			continue
		}
		// Start with the parent's TextStat. Then ascend and check if the
		// current chunk has an ancestor with better stats. Use the best stat
		// as result.
		stat := ancestorStat[node]
		for {
			if node = node.Parent; node == nil {
				break
			}
			if statPrev, ok := ancestorStat[node]; ok {
				if stat.Count < statPrev.Count {
					stat = statPrev
				}
			} else {
				break
			}
		}
		result[chunk] = stat
	}
	return result
}
