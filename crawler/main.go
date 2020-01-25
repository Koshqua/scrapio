package crawler

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//Package crawler defines all the functionality for page crawling

//Crawler defines a default crawler
type Crawler struct {
	BaseURL  string
	StartURL string
	Results  []CrawlResult
}

//CrawlResults ...
type CrawlResult struct {
	URL   string
	Title string
}

//GetRequest ...
func (c *Crawler) GetRequest(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromResponse(res)
	return doc, nil
}

//ParseBase ...
func (c *Crawler) ParseBase() error {
	url, err := url.Parse(c.StartURL)
	if err != nil {
		return err
	}
	c.BaseURL = fmt.Sprintf("%s://%s", url.Scheme, url.Host)
	return nil
}

func (c *Crawler) FormatRelative(urls map[string]int) (formatedUrls []string) {
	c.ParseBase()
	for url, _ := range urls {
		if strings.HasPrefix(url, c.BaseURL) {
			formatedUrls = append(formatedUrls, url)
		}
		if strings.HasPrefix(url, "/") {
			formated := fmt.Sprintf("%s%s", c.BaseURL, url)
			formatedUrls = append(formatedUrls, formated)
		}
	}
	return formatedUrls
}

//GetLinks..
func (c *Crawler) GetLinks(doc *goquery.Document) []string {
	foundLinks := make(map[string]int)
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		if _, ok := foundLinks[link]; !ok {
			foundLinks[link] = 1
		}
	})
	return c.FormatRelative(foundLinks)
}

func (c *Crawler) GetResult(doc *goquery.Document, url string) {
	cr := CrawlResult{}
	cr.URL = url
	cr.Title = doc.Find("title").Text()
	c.Results = append(c.Results, cr)
}

func (c *Crawler) CrawlPage(url string) ([]string, error) {
	doc, err := c.GetRequest(url)
	if err != nil {
		return nil, err
	}
	c.GetResult(doc, url)
	links := c.GetLinks(doc)
	return links, nil
}
