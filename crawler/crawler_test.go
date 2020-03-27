package crawler

import (
	"testing"
)

func BenchmarkCrawl(b *testing.B) {
	b.ReportAllocs()
	cr := Crawler{StartURL: "https://blog.merovius.de/"}
	cr.Crawl()
}

// func BenchmarkCrawlPage(b *testing.B) {
// 	b.ReportAllocs()
// 	startingURL := "https://blog.merovius.de/"
// 	cr := Crawler{StartURL: "https://blog.merovius.de/"}
// 	links, crawlResult, err := cr.CrawlPage(startingURL)
// 	if err != nil {
// 		b.Errorf("%v", err.Error())
// 	}
// 	if crawlResult.URL != startingURL {
// 		b.Errorf("Different URLS")
// 	}
// 	if len(links) == 0 {
// 		b.Errorf("No links were found")
// 	}
// }
