package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/koshqua/scrapio/crawler"
	"github.com/koshqua/scrapio/scraper"
)

func main() {
	start := time.Now()
	cr := &crawler.Crawler{StartURL: "https://gulfnews.com/"}
	startCrawling := time.Now()
	cr.Crawl(2000)
	finishCrawling := time.Now()
	h2 := scraper.NewSelector("h2", true, true, true)
	img := scraper.NewSelector("img", true, true, true)
	p := scraper.NewSelector("p:first-of-type", true, true, true)
	sc := scraper.InitScraper(*cr, []scraper.Selector{h2, img, p})
	startScrap := time.Now()
	err := sc.Scrap()
	finishScrap := time.Now()
	finish := time.Now()

	if err != nil {
		log.Fatalln(err)
	}
	js, err := json.MarshalIndent(*sc, "", " ")
	if err != nil {
		log.Fatalln(err)
	}
	file, err := os.Create("result.json")
	if err != nil {
		log.Fatalln(err)
	}
	file.Write(js)

	fmt.Println("Time for crawling: ", finishCrawling.Sub(startCrawling))
	fmt.Println("Time for scraping: ", finishScrap.Sub(startScrap))
	fmt.Println(finish.Sub(start))

}
