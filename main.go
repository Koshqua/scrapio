package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/koshqua/scrapio/crawler"
	"github.com/koshqua/scrapio/scraper"
)

func main() {
	cr := &crawler.Crawler{StartURL: "https://blog.merovius.de/"}
	cr.PagesLimit = 11
	cr.Crawl()
	h2 := scraper.NewSelector("p:nth-of-type(2)", true, true, true)
	img := scraper.NewSelector("figure img", true, true, true)
	sc := scraper.InitScraper(*cr, []scraper.Selector{h2, img})
	err := sc.Scrap()
	if err != nil {
		log.Fatalln(err)
	}
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

}
