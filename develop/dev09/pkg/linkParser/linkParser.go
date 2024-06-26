// link is a package for parsing HTML link tags (<a href="..."...</a>),
package linkParser

import (
	"io"
	"net/url"

	"golang.org/x/net/html"
)

// Link represents HTML link tag
type Link struct {
	Href *url.URL
}

// ParseHTML parses given html file and returns slice of links
func ParseHTML(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	links := parseNode(doc)

	return links, nil
}

// parseNode gets links from html Body
func parseNode(node *html.Node) []Link {
	links := make([]Link, 0)

	if node.Type == html.ElementNode && node.Data == "a" {
		u, err := parseHref(node.Attr)
		if err != nil {
			return links
		}
		link := Link{
			Href: u,
		}

		links = append(links, link)
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		parseNode(c)
	}

	return links
}

// parseHref extracts href attribute from link tag
func parseHref(attrs []html.Attribute) (*url.URL, error) {
	var href string

	for _, a := range attrs {
		if a.Key == "href" {
			href = a.Val
			break
		}
	}
	u, err := url.Parse(href)
	if err != nil {
		return nil, err
	}

	return u, nil
}
