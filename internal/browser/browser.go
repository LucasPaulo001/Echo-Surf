package browser

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

type Page struct {
	URL 		string
	StatusCode	int
	Body		[]byte
	Title 		string
	LinkCount 	int
	ImagesCount int
	Headers 	http.Header
}

func LoadPage(url string) (*Page, error) {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := client.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	title, links, imgs := Extract(bytes.NewReader(body))

	page := &Page{
		URL: 		 url,	
		StatusCode:  resp.StatusCode,
		Body: 		 body,	
		Title: 		 title,	
		LinkCount:   links,
		ImagesCount: imgs,
		Headers:     resp.Header,
	}

	return page, nil
}
