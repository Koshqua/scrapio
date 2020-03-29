package crawler

import (
	"encoding/json"
	"log"
	"os"
	"testing"
)

func ExampleCrawl() {
	c := Crawler{StartURL: "https://blog.merovius.de/"}
	c.Crawl()
	json, err := json.Marshal(c)
	if err != nil {
		log.Fatalln(err)
	}
	f, err := os.Create("result.josn")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = f.Write(json)
	if err != nil {
		log.Fatalln(err)
	}
}

func BenchmarkCrawl(b *testing.B) {
	b.ReportAllocs()
	cr := Crawler{StartURL: "https://blog.merovius.de/"}
	cr.Crawl()
}

func BenchmarkCrawlPage(b *testing.B) {
	b.ReportAllocs()
	startingURL := "https://blog.merovius.de/"
	cr := Crawler{StartURL: "https://blog.merovius.de/"}
	links, crawlResult, err := cr.CrawlPage(startingURL)
	if err != nil {
		b.Errorf("%v", err.Error())
	}
	if crawlResult.URL != startingURL {
		b.Errorf("Different URLS")
	}
	if len(links) == 0 {
		b.Errorf("No links were found")
	}
}
