package browser

import (
	"io"
	"strings"
	"golang.org/x/net/html"
)

func Extract(r io.Reader) (string, []string, []string) {
	doc, err := html.Parse(r)

	if err != nil {
		return "", nil, nil
	}

	var title string
	var links []string
	var images  []string

	var crawl func(*html.Node)
	crawl = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "title":
				if n.FirstChild != nil {
					title = strings.TrimSpace(n.FirstChild.Data)
				}
			case "a":
				for _, attr := range n.Attr {
					if attr.Key == "href" && attr.Val != "" {
						links = append(links, attr.Val)
					}
				}
			case "img":
				for _, attr := range n.Attr {
					if attr.Key == "src" && attr.Val != "" {
						images = append(images, attr.Val)
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			crawl(c)
		}
	}

	crawl(doc)
	return title, links, images
}