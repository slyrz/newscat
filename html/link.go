package html

import (
	"code.google.com/p/go.net/html"
	"errors"
	"net/url"
)

// TODO
type Link struct {
	URL *url.URL
}

// TODO
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

func (l *Link) Resolve(base *url.URL) {
	l.URL = base.ResolveReference(l.URL)
}
