package server

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/koshqua/scrapio/crawler"
	"github.com/koshqua/scrapio/utils"
)

//Register handles user registration
type Register struct {
	db *sql.DB
}

//Register handles register route
func (reg Register) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	utils.CheckParse(err, w, "Bad request", http.StatusBadRequest)

}

//CrawlHandler ...
type CrawlHandler struct {
	db *sql.DB
}

//CrawlHandler ...
func (ch CrawlHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	utils.CheckParse(err, w, "Bad request", http.StatusBadRequest)
	ct := r.Header.Get("Content-Type")
	crawler := new(crawler.Crawler)
	switch {
	case ct == "application/json":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		json.Unmarshal(body, crawler)
	case ct == "application/x-www-form-urlencoded":
		url := r.FormValue("StartURL")
		crawler.StartURL = url
	}
	crawler.Crawl()
	w.Header().Set("Content-Type", "application/json")
	bs, err := json.Marshal(crawler.Results)
	w.Write(bs)
}
