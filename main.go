package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/koshqua/scrapio/crawler"
	"github.com/koshqua/scrapio/scraper"
)

func main() {
	cr := &crawler.Crawler{StartURL: "https://adnanahmed.info/"}
	cr.Crawl()
	selectors := scraper.Selector{
		"figure img",
		true,
		true,
		true,
		scraper.ScrapResult{},
	}
	sc := scraper.InitScraper(*cr, []scraper.Selector{selectors})
	err := sc.Scrap()
	if err != nil {
		log.Fatalln(err)
	}
	js, err := json.Marshal(*sc)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(js))

}
