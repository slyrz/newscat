package html

import (
	"code.google.com/p/go.net/html"
	"errors"
	"github.com/slyrz/newscat/util"
	"io"
)

// Errors returned during Document parsing.
var (
	ErrNoHTML = errors.New("missing html element")
	ErrNoHead = errors.New("missing head element")
	ErrNoBody = errors.New("missing body element")
)

// Document is a parsed HTML document that extracts the document title and
// holds unexported pointers to the html, head and body nodes.
type Document struct {
	Title *util.Text // the <title>...</title> text.

	// Unexported fields.
	html *html.Node // the <html>...</html> part
	head *html.Node // the <head>...</head> part
	body *html.Node // the <body>...</body> part
}

// NewDocument parses the HTML data provided through an io.Reader interface.
func NewDocument(r io.Reader) (*Document, error) {
	doc := new(Document)
	if err := doc.init(r); err != nil {
		return nil, err
	}
	return doc, nil
}

func (doc *Document) init(r io.Reader) error {
	doc.Title = util.NewText()

	root, err := html.Parse(r)
	if err != nil {
		return err
	}

	// Assign the fields html, head and body from the HTML page.
	iterateNode(root, func(n *html.Node) int {
		switch n.Data {
		case "html":
			doc.html = n
			return IterNext
		case "body":
			doc.body = n
			return IterSkip
		case "head":
			doc.head = n
			return IterSkip
		}
		// Keep going as long as we're missing some nodes.
		return IterNext
	})

	switch {
	case doc.html == nil:
		return ErrNoHTML
	case doc.head == nil:
		return ErrNoHead
	case doc.body == nil:
		return ErrNoBody
	}

	// Detect the document title.
	iterateNode(doc.head, func(n *html.Node) int {
		if n.Type == html.ElementNode && n.Data == "title" {
			iterateText(n, doc.Title.WriteString)
			return IterStop
		}
		return IterNext
	})
	return nil
}
