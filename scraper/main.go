package scraper

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/koshqua/scrapio/crawler"
)

//Scraper represents default scraper
type Scraper struct {
	StartURL string
	BaseURL  string
	Pages    map[string]map[string][]string
}

//InitScraper creates a basic scraper from Crawler
func (s *Scraper) InitScraper(c crawler.Crawler) {
	s.BaseURL = c.BaseURL
	s.StartURL = c.StartURL
	for _, result := range c.Results {
		s.Pages[result.URL] = map[string][]string{}
	}
}

//ScrapPageText scraps single page ...
func ScrapPageText(selectors []string, url string) (map[string][]string, error) {
	results := make(map[string][]string)
	res, err := http.Get(url)
	if err != nil {
		return results, err
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return results, err
	}
	defer res.Body.Close()
	for _, selector := range selectors {
		doc.Find(selector).Each(func(i int, s *goquery.Selection) {
			content := s.Text()
			results[selector] = append(results[selector], content)
		})
	}
	return results, nil
}

func parseCrawler(j []byte) (crawler.Crawler, error) {
	c := crawler.Crawler{}
	err := json.Unmarshal(j, &c)
	if err != nil {
		return c, err
	}
	return c, nil
}

func parseSelectors(s string) []string {
	var selArr []string
	selArr = strings.Split(s, ", ")
	return selArr
}
