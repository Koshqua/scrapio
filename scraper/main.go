package scraper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/koshqua/scrapio/crawler"
)

//ScrapResult ...
type ScrapResult struct {
	Text     string
	LinkURL  string
	ImageURL string
}

//Selector ...
type Selector struct {
	Name        string `json:"Name"`
	ScrapText   bool   `json:"ScrapText"`
	ScrapLinks  bool   `json:"ScrapLinks"`
	ScrapImages bool   `json:"ScrapImages"`
	ScrapResult
}

//Page ...
type Page struct {
	URL       string
	Selectors []*Selector
}

//Scraper represents default scraper
type Scraper struct {
	ID      string
	BaseURL string
	Pages   []*Page
}

//parseScraper creates a basic scraper from Crawler
func parseScraper(c crawler.Crawler) *Scraper {
	s := &Scraper{}
	s.BaseURL = c.BaseURL
	for _, result := range c.Results {
		s.Pages = append(s.Pages, &Page{URL: result.URL})
	}
	return s
}

func (s *Scraper) addSelectors(selectors []Selector) {
	for _, page := range s.Pages {
		for _, s := range selectors {
			page.Selectors = append(page.Selectors, &s)
		}
	}
}

//InitScraper creates a Scraper with selectors attached to it
func InitScraper(c crawler.Crawler, s []Selector) *Scraper {
	scraper := parseScraper(c)
	scraper.addSelectors(s)
	return scraper
}

func (s *Scraper) Scrap() error {
	for _, page := range s.Pages {
		err := scrapPage(page)
		if err != nil {
			return err
		}
	}
	return nil
}
func scrapPage(p *Page) error {
	res, err := http.Get(p.URL)
	if err != nil {
		return err
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	for _, selector := range p.Selectors {
		if selector.ScrapImages {
			scrapPageImage(doc, selector)
		}
		if selector.ScrapLinks {
			scrapPageLinks(doc, selector)
		}
		if selector.ScrapText {
			scrapPageText(doc, selector)
		}
	}
	return nil
}

//ScrapPageText scraps single page ...
func scrapPageText(doc *goquery.Document, selector *Selector) {
	selection := doc.Find(selector.Name).First()
	selector.Text = selection.Text()
}

func scrapPageImage(doc *goquery.Document, selector *Selector) {
	selection := doc.Find(selector.Name).First()
	url, _ := selection.Attr("src")
	fmt.Println(url)
	selector.ImageURL, _ = selection.Attr("src")
}
func scrapPageLinks(doc *goquery.Document, selector *Selector) {
	selection := doc.Find(selector.Name).First()
	selector.LinkURL, _ = selection.Attr("href")

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
