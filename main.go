package main

import (
	"fmt"
	"net/http"

	"github.com/koshqua/scrapio/server"
)

func main() {
	http.Handle("/crawl", server.CrawlHandler{})
	fmt.Println("Server is started")
	http.ListenAndServe(":3000", nil)
}
