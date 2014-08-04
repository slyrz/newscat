package html

import (
	"code.google.com/p/go.net/html"
	"github.com/slyrz/newscat/util"
	"io"
	"unicode"
)

const (
	chunkCap = 512 // initial capacity of the Article.Chunks array
)

const (
	// We remember a few special node types when descending into their
	// children.
	AncestorArticle = 1 << iota
	AncestorAside
	AncestorBlockquote
	AncestorList
)

// Article stores all text chunks found in a HTML document.
type Article struct {
	Document
	Chunks []*Chunk // all chunks found in this document.

	// Unexported fields used as state variables during parsing.
	ancestors int                // bitmask to track specific ancestor types
	linkText  map[*html.Node]int // length of text inside <a></a> tags
	normText  map[*html.Node]int // length of text outside <a></a> tags
}

// NewArticle parses the HTML data provided through an io.Reader interface
// and returns, if successful, an Article object that can be used to access
// all relevant text chunks found in the document.
func NewArticle(r io.Reader) (*Article, error) {
	article := new(Article)
	if err := article.init(r); err != nil {
		return nil, err
	}
	return article, nil
}

func (article *Article) init(r io.Reader) error {
	if err := article.Document.init(r); err != nil {
		return err
	}

	article.Chunks = make([]*Chunk, 0, chunkCap)
	article.linkText = make(map[*html.Node]int)
	article.normText = make(map[*html.Node]int)

	article.cleanBody(article.body, 0)
	article.countText(article.body, false)
	article.parseBody(article.body)

	// Now we link the chunks.
	min, max := 0, len(article.Chunks)-1
	for i := range article.Chunks {
		if i > min {
			article.Chunks[i].Prev = article.Chunks[i-1]
		}
		if i < max {
			article.Chunks[i].Next = article.Chunks[i+1]
		}
	}
	return nil
}

// countText counts the text inside of links and the text outside of links
// per html.Node. Counting is done cumulative, so the umbers of a parent node
// include the numbers of it's child nodes.
func (article *Article) countText(n *html.Node, insideLink bool) (linkText int, normText int) {
	linkText = 0
	normText = 0
	if n.Type == html.ElementNode && n.Data == "a" {
		insideLink = true
	}
	for s := n.FirstChild; s != nil; s = s.NextSibling {
		linkTextChild, normTextChild := article.countText(s, insideLink)
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
	article.linkText[n] = linkText
	article.normText[n] = normText
	return
}

// cleanBody removes unwanted HTML elements from the HTML body.
func (article *Article) cleanBody(n *html.Node, level int) {

	// removeNode returns true if a node should be removed from HTML document.
	removeNode := func(c *html.Node, level int) bool {
		switch c.Data {
		// Elements save to ignore.
		case "address", "audio", "button", "canvas", "caption", "fieldset",
			"figcaption", "figure", "footer", "form", "frame", "iframe",
			"map", "menu", "nav", "noscript", "object", "option", "output",
			"script", "select", "style", "svg", "textarea", "video":
			return true
		// High-level tables might be used to layout the document, so we better
		// not ignore them.
		case "table":
			return level > 5
		}
		return false
	}

	var curr *html.Node = n.FirstChild
	var next *html.Node = nil
	for ; curr != nil; curr = next {
		// We have to remember the next sibling here because calling RemoveChild
		// sets curr's NextSibling pointer to nil and we would quit the loop
		// prematurely.
		next = curr.NextSibling
		if curr.Type == html.ElementNode {
			if removeNode(curr, level) {
				n.RemoveChild(curr)
			} else {
				article.cleanBody(curr, level+1)
			}
		}
	}
}

var (
	ignoreNames = util.NewRegexFromWords(
		"breadcrumb",
		"byline",
		"caption",
		"comment",
		"community",
		"credit",
		"description",
		"email",
		"foot",
		"gallery",
		"hide",
		"infotext",
		"photo",
		"related",
		"shares",
		"social",
		"story[-_]?bar",
		"story[-_]?feature",
	)
	ignoreStyle = util.NewRegex(`(?i)display:\s*none`)
)

// parseBody parses the <body>...</body> part of the HTML page. It creates
// Chunks for every html.TextNode found in the body.
func (article *Article) parseBody(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		// We ignore the node if it has some nasty classes/ids/itemprops or if
		// its style attribute contains "display: none".
		for _, attr := range n.Attr {
			switch attr.Key {
			case "id", "class", "itemprop":
				if ignoreNames.In(attr.Val) {
					return
				}
			case "style":
				if ignoreStyle.In(attr.Val) {
					return
				}
			}
		}
		ancestorMask := 0
		switch n.Data {
		// We convert headings and links to text immediately. This is easier
		// and feasible because headings and links don't contain many children.
		// Descending into these children and handling every TextNode separately
		// would make things unnecessary complicated and our results noisy.
		case "h1", "h2", "h3", "h4", "h5", "h6", "a":
			if chunk, err := NewChunk(article, n); err == nil {
				article.Chunks = append(article.Chunks, chunk)
			}
			return
		// Now mask the element type, but only if it isn't already set.
		// If we mask a bit which was already set by one of our callers, we'd also
		// clear it at the end of this function, though it actually should be cleared
		// by the caller.
		case "article":
			ancestorMask = AncestorArticle &^ article.ancestors
		case "aside":
			ancestorMask = AncestorAside &^ article.ancestors
		case "blockquote":
			ancestorMask = AncestorBlockquote &^ article.ancestors
		case "ul", "ol":
			ancestorMask = AncestorList &^ article.ancestors
		}
		// Add our mask to the ancestor bitmask.
		article.ancestors |= ancestorMask
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			article.parseBody(c)
		}
		// Remove our mask from the ancestor bitmask.
		article.ancestors &^= ancestorMask
	case html.TextNode:
		if chunk, err := NewChunk(article, n); err == nil {
			article.Chunks = append(article.Chunks, chunk)
		}
	}
}

// TextStat contains the number of words and sentences found in text.
type TextStat struct {
	Words     int // total number of words
	Sentences int // total number of sentences
	Count     int // number of texts used to calculate this stats
}

// GetClassStats groups the document chunks by their classes (defined by the
// class attribute of HTML nodes) and calculates TextStats for each class.
func (article *Article) GetClassStats() map[string]*TextStat {
	result := make(map[string]*TextStat)
	for _, chunk := range article.Chunks {
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
func (article *Article) GetClusterStats() map[*Chunk]*TextStat {
	// Don't ascend further than this.
	const maxAncestors = 3

	// Count TextStats for Chunk ancestors.
	ancestorStat := make(map[*html.Node]*TextStat)
	for _, chunk := range article.Chunks {
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

	// Generate result. For each chunk pick the best TextStats from its ancestors.
	result := make(map[*Chunk]*TextStat)
	for _, chunk := range article.Chunks {
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
