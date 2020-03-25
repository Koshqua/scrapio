package crawler

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

//Package crawler defines all the functionality for page crawling

//Crawler defines a default crawler
type Crawler struct {
	ID       string        `json:"ID"`
	BaseURL  string        `json:"BaseURL"`
	StartURL string        `json:"StartURL"`
	Results  []CrawlResult `json:"Results"`
}

//CrawlResult ...
type CrawlResult struct {
	URL   string `json:"URL"`
	Title string `json:"Title"`
}

//GetRequest ...
func (c *Crawler) GetRequest(url string) (*goquery.Document, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html")
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	html, err := html.Parse(res.Body)
	if err != nil {
		return nil, err
	}
	doc := goquery.NewDocumentFromNode(html)
	res.Body.Close()
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

//FormatRelative ...
func (c *Crawler) FormatRelative(urls map[string]int) (formatedUrls []string) {
	if c.BaseURL == "" {
		c.ParseBase()
	}
	for url := range urls {
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

//GetLinks ...
func (c *Crawler) GetLinks(doc *goquery.Document) []string {

	foundLinks := make(map[string]int)
	if doc != nil {
		doc.Find("a").Each(func(i int, s *goquery.Selection) {
			link, _ := s.Attr("href")
			if _, ok := foundLinks[link]; !ok {
				foundLinks[link] = 1
			}
		})
	}

	return c.FormatRelative(foundLinks)
}

//GetResult ...
func (c *Crawler) GetResult(doc *goquery.Document, url string) CrawlResult {
	cr := CrawlResult{}
	cr.URL = url
	if doc != nil {
		cr.Title = doc.Find("title").Text()
	}
	return cr
}

//CrawlPage ...
func (c *Crawler) CrawlPage(url string) ([]string, CrawlResult, error) {
	doc, err := c.GetRequest(url)
	if err != nil {
		return nil, CrawlResult{}, err
	}
	cr := c.GetResult(doc, url)
	links := c.GetLinks(doc)
	return links, cr, nil
}

//Crawl ....
func (c *Crawler) Crawl() {
	results := make(chan CrawlResult)
	worklist := make(chan []string)
	seen := make(map[string]bool)
	go func() { worklist <- []string{c.StartURL} }()
	go func() {
		for {
			cr := <-results
			c.Results = append(c.Results, cr)
		}
	}()

	for n := 1; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if _, ok := seen[link]; !ok {
				seen[link] = true
				n++
				go func(link string) {
					links, cr, err := c.CrawlPage(link)
					if err != nil {
						fmt.Println(err)
						return
					}
					fmt.Println(cr)
					worklist <- links
					results <- cr
				}(link)

			}
		}
	}

}
