package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/koshqua/scrapio/crawler"
)

var tmp *template.Template

func init() {
	var err error
	tmp, err = tmp.ParseGlob("./assets/*.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

}

func main() {
	router := httprouter.New()
	router.ServeFiles("/assets/*filepath", http.Dir("./assets"))
	router.GET("/", Index)
	router.POST("/crawl", CrawlHandler)
	http.ListenAndServe(":3000", router)

}
func CrawlHandler(res http.ResponseWriter, req *http.Request, param httprouter.Params) {
	url := req.FormValue("start_url")
	limit, err := strconv.Atoi(req.FormValue("limit"))
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}
	c := crawler.Crawler{StartURL: url}
	c.PagesLimit = limit
	err = c.Crawl()
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}
	json, err := json.Marshal(c)
	if err != nil {
	}
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(200)

	res.Write(json)
}

func Index(res http.ResponseWriter, req *http.Request, param httprouter.Params) {
	err := tmp.ExecuteTemplate(res, "index.gohtml", nil)
	if err != nil {
		http.Error(res, "impossible to execute template", http.StatusBadRequest)
	}
}
