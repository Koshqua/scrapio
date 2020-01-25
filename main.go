package main

import (
	"fmt"

	"github.com/koshqua/scrapio/crawler"
)

func main() {
	c := crawler.Crawler{StartURL: "https://edmundmartin.com/"}
	l, err := c.CrawlPage(c.StartURL)
	if err != nil {
		fmt.Println(err)
	}
	for _, link := range l {
		fmt.Println(link)
	}
	fmt.Println(c.Results)
}
