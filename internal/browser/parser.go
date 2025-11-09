package browser

import (
	"io"

	"github.com/PuerkitoBio/goquery"
)

func Extract(r io.Reader) (string, int, int) {
	doc, err := goquery.NewDocumentFromReader(r)

	if err != nil {
		return "", 0, 0
	}

	title := doc.Find("title").First().Text()
	links := doc.Find("a").Length()
	imgs := doc.Find("img").Length()

	return title, links, imgs
}