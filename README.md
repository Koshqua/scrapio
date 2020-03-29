
[![GoDoc](https://godoc.org/github.com/koshqua/scrapio?status.svg)](https://pkg.go.dev/github.com/koshqua/scrapio)

## Scrapio 

**Scrapio** - is a lightweight and user-friendy web crawling and scraping library. 
The main goal of creating the project was to make scraping big amounts of similar data from web easy and user-friendly. It might be useful for wide range of applications, like data mining, data processing and archiving. 
After some time, I am going to make it a standalone service, which will work as an API.




### Features
At the moment works as a library which can be used to crawl and scrap data from web. 
What it can do:
- Crawl all pages on host, return all the links. 
- Scrap text, image urls and links from Crawl Result pages. 
- It leaves the choice of data output(csv,json, etc) up to you. 
- It's free and quite powerful. 
- Written in go, concurrent, depending on Network Speed can crawl and scrap up to 2k pages/minute.

### Usage 
**Crawler** is easy to use. You just need to specify a starting URL and it will crawl all the URL on the host. 
```go 
    //init a new crawler, give it a start url, it's not necessary should be basic URL
    cr := &crawler.Crawler{StartURL: "https://gulfnews.com/"}
    //Start crawling func. 
    //After some time im going to implement more configs for this func, like max results, etc.
    cr.Crawl()
    //Do something with result, it's up to you
```
**Scraper** uses data structure given by crawler. 
Before initiating a scraper, you need to create a few selectors, to assign them to scraper.
Selectors are the simple css-like selectors.  
```go
    //create some Selectors, which you want to scrap.
    h2 := scraper.NewSelector("h2", true, true, true)
    img := scraper.NewSelector("img", true, true, true)
    p := scraper.NewSelector("p:first-of-type", true, true, true)
    //Initiate a new scrapper with given selectors
    //Scraper depends on the crawler from previous code snippet.
    //It gets pages and creates new structure with selectors and scrap results.
    sc := scraper.InitScraper(*cr, []scraper.Selector{h2, img, p})
    //And just start scraping
	err := sc.Scrap()
	if err != nil {
		log.Fatalln(err)
	}
```


