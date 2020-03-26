### Scrapio 
**Scrapio** - is a lightweight and user-friendy web crawling and scraping library. 
The main goal of creating the project was to make scraping big amounts of similar data from web easy and user-friendly.
After some time, I am going to make it a standalone service, which will work as an API.
BTW, at the moment I'm looking for a front-end developer, who'd like to commit.
If you are the one, you can email me on maliovaniy.ivan@gmail.com


### Progress and some user guide.
At the moment works as a library which can be used to crawl and scrap data from web. 
What it can do:
- Crawl all pages on host, return all the links. 
- Scrap text, image urls and links from Crawl Result pages. 
- It leaves the choice of data output(csv,json, etc) up to you. 
- It's free and quite powerful. 
- Written in go, concurrent, depending on Network Speed can crawl and scrap up to 2k pages/minute.

### Usage 
- Crawler 
```go 
    //init a new crawler, give it a start url, it's not necessary should be basic URL
    cr := &crawler.Crawler{StartURL: "https://gulfnews.com/"}
    //Start crawling func. 
    //After some time im going to implement more configs for this func, like max results, etc.
    cr.Crawl()
    //Do something with result, it's up to you
```
- Scraper 
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

### FEATURES ###
- Should work as an api
- Should be able to scrap data from web pages by link or by a list of links. 
- Should scrap data from elements by css selectors. 
- Should parse a json or csv table with results (preferably json).
- Should have an authentication.

