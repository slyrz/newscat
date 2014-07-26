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

// NewLink creates a new link from a HTML anchor node.
func NewLink(n *html.Node) (*Link, error) {
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			if url, err := url.Parse(attr.Val); err != nil {
				return nil, err
			} else {
				return &Link{url}, nil
			}
		}
	}
	return nil, errors.New("href not found")
}

// Resolves joins link with an absoulte base URL if link isn't absolute yet.
func (l *Link) Resolve(base *url.URL) {
	if l.URL.Scheme == "" || l.URL.Host == "" {
		l.URL = base.ResolveReference(l.URL)
	}
}
