package html

import (
	"code.google.com/p/go.net/html"
	"io"
	"net/url"
)

const (
	linkCap = 256 // initial capacity of the Website.Links array
	feedCap = 4   // initial capacity of the Website.Feeds array
)

// Website finds all links in a HTML document.
type Website struct {
	Document
	Links []*Link // all links found in this document.
	Feeds []*Link // all RSS feeds found in this document.
}

// NewWebsite parses the HTML data provided through an io.Reader interface
// and returns, if successful, a Website object that can be used to access
// all links and extract links to news articles.
func NewWebsite(r io.Reader) (*Website, error) {
	website := new(Website)
	if err := website.init(r); err != nil {
		return nil, err
	}
	return website, nil
}

func (website *Website) init(r io.Reader) error {
	if err := website.Document.init(r); err != nil {
		return err
	}
	website.Links = make([]*Link, 0, linkCap)
	website.Feeds = make([]*Link, 0, feedCap)

	// Extract all links.
	iterateNode(website.body, func(n *html.Node) int {
		if n.Type == html.ElementNode && n.Data == "a" {
			if link, err := NewLink(n); err == nil {
				website.Links = append(website.Links, link)
			}
			return IterSkip
		}
		return IterNext
	})

	// Extract all RSS feeds.
	iterateNode(website.head, func(n *html.Node) int {
		if n.Data != "link" {
			return IterNext
		}
		// Scan the link attributes and make sure we find
		//	rel="alternate" type="application/rss+xml" href="..."
		href, hasRel, hasType := "", false, false
		for _, attr := range n.Attr {
			switch attr.Key {
			case "rel":
				hasRel = attr.Val == "alternate"
			case "type":
				hasType = attr.Val == "application/rss+xml"
			case "href":
				href = attr.Val
			}
		}
		if hasRel && hasType && href != "" {
			if link, err := NewLinkFromString(href); err == nil {
				website.Feeds = append(website.Feeds, link)
			}
		}
		return IterNext
	})
	return nil
}

// ResolveBase transforms relative URLs to absolute URLs by adding
// missing components from an absolute base URL.
func (website *Website) ResolveBase(base string) error {
	baseURL, err := url.Parse(base)
	if err == nil {
		for _, link := range website.Links {
			link.Resolve(baseURL)
		}
		for _, feed := range website.Feeds {
			feed.Resolve(baseURL)
		}
	}
	return err
}
