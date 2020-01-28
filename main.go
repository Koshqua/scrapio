package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/koshqua/scrapio/crawler"
)

func main() {
	t := time.Now()
	c := crawler.Crawler{StartURL: "http://www.greendeco.com.ua/"}
	c.Crawl()
	url, err := url.Parse(c.StartURL)

	f, err := os.Create(strings.TrimPrefix(url.Hostname(), "www.") + ".json")
	if err != nil {
		log.Fatalln(err)
	}
	enc := json.NewEncoder(f)
	enc.SetIndent("", " ")
	err = enc.Encode(c.Results)
	t2 := time.Now()
	diff := t2.Sub(t)
	fmt.Println(diff)
}
