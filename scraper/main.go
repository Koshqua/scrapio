package scraper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

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
			s := s
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
	wg := sync.WaitGroup{}
	pageChan := make(chan *Page, len(s.Pages))
	resultChan := make(chan *Page, len(s.Pages))
	for _, page := range s.Pages {
		page := page
		fmt.Println("Transfering page", page.URL)
		pageChan <- page
	}

	counter := len(s.Pages)
	s.Pages = []*Page{}
	fmt.Println("Cleared all the pages from struct")

	go func() {
		wg.Add(1)
		for n := 0; n < counter; {
			page := <-resultChan
			fmt.Println("Put result page", page.URL)
			s.Pages = append(s.Pages, page)
			n++
		}
		wg.Done()
	}()
	for n := 0; n < counter; {
		page := <-pageChan
		n++
		fmt.Println("Scraping loop", n)
		go func() {
			page := page
			fmt.Println(page.URL, "Got from chanel")
			fmt.Println("Scraping page", page.URL)
			p, err := scrapPage(page)
			if err != nil {
				log.Fatalln(err)
				return
			}
			resultChan <- p

		}()

	}
	wg.Wait()

	return nil
}
func scrapPage(p *Page) (*Page, error) {
	page := p
	fmt.Println("MADE REQ TO ", page.URL)
	res, err := http.Get(page.URL)
	fmt.Println("FINISHED REQ TO", page.URL)
	if err != nil {
		return page, err
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		return page, err
	}
	defer res.Body.Close()
	for _, selector := range page.Selectors {
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
	fmt.Println("scrapped page", page.URL)
	return page, nil
}

//ScrapPageText scraps single page ...
func scrapPageText(doc *goquery.Document, selector *Selector) {
	selection := doc.Find(selector.Name).First()
	selector.Text = selection.Text()
}

func scrapPageImage(doc *goquery.Document, selector *Selector) {
	selection := doc.Find(selector.Name).First()
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

func NewSelector(name string, scrapPageImages bool, scrapPageLinks bool, scrapPageText bool) Selector {
	s := Selector{
		Name:        name,
		ScrapImages: scrapPageImages,
		ScrapLinks:  scrapPageLinks,
		ScrapText:   scrapPageText,
	}
	return s
}
