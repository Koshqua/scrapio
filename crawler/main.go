//Package crawler defines all the functionality for page crawling
package crawler

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//Crawler defines a default crawler
type Crawler struct {
	ID       string        `json:"ID"`
	BaseURL  string        `json:"BaseURL"`
	StartURL string        `json:"StartURL"`
	Results  []CrawlResult `json:"Results"`
}

//CrawlResult defines the result of crawled single page.
type CrawlResult struct {
	URL   string `json:"URL"`
	Title string `json:"Title"`
}

var errTimeOut = errors.New("Connection timed out")

//Crawl crawls the whole host of give startURL and saves data(URLs and Titles) to Crawler struct.
func (c *Crawler) Crawl() error {
	errChan := make(chan error)
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
				select {
				case <-errChan:
					err := <-errChan
					if err == errTimeOut {
						return err
					}
				default:
					seen[link] = true
					n++
					go func(link string) {
						links, cr, err := c.CrawlPage(link)
						if err != nil {
							errChan <- err
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
	return nil
}

//CrawlPage crawls single page, returns links as []string, CrawlResult(Page URL and Title) and error.
func (c *Crawler) CrawlPage(url string) ([]string, CrawlResult, error) {
	doc, err := c.GetRequest(url)
	if err != nil {
		return nil, CrawlResult{}, err
	}
	cr := c.GetResult(doc, url)
	links := c.GetLinks(doc)
	return links, cr, nil
}

//GetRequest is a helper function for CrawlPage.
//It makes a request to a page and returns goquery.Document and error.
func (c *Crawler) GetRequest(url string) (*goquery.Document, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html")
	client := http.Client{}
	res, err := client.Do(req)
	if res.StatusCode == http.StatusGatewayTimeout || res.StatusCode == http.StatusRequestTimeout {
		return nil, errTimeOut
	} else if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil

}

//GetLinks return all the a[href] values from the goquery.Document
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

//GetResult returns a CrawlResult from a single page
func (c *Crawler) GetResult(doc *goquery.Document, url string) CrawlResult {
	cr := CrawlResult{}
	cr.URL = url
	if doc != nil {
		cr.Title = doc.Find("title").Text()
	}
	return cr
}

//FormatRelative formats relative links to an absolute links if encounter them during crawling.
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

//ParseBase parses basic url of host.
func (c *Crawler) ParseBase() error {
	url, err := url.Parse(c.StartURL)
	if err != nil {
		return err
	}
	c.BaseURL = fmt.Sprintf("%s://%s", url.Scheme, url.Host)
	return nil
}
