package html

import (
	"code.google.com/p/go.net/html"
	"errors"
	"net/url"
)

// Link holds a URL.
type Link struct {
	URL *url.URL
}

// NewLinkFromString creates a new link by parsing a URL string.
func NewLinkFromString(s string) (*Link, error) {
	if url, err := url.Parse(s); err == nil {
		return &Link{url}, nil
	} else {
		return nil, err
	}
}

// NewLink creates a new link from a HTML anchor node.
func NewLink(n *html.Node) (*Link, error) {
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			return NewLinkFromString(attr.Val)
		}
	}
	return nil, errors.New("href not found")
}

// Resolves joins link with an absolute base URL if link isn't absolute yet.
func (l *Link) Resolve(base *url.URL) {
	if l.URL.Scheme == "" || l.URL.Host == "" {
		l.URL = base.ResolveReference(l.URL)
	}
}
