package server

import (
	"database/sql"
	"github.com/koshqua/scrapio/utils"
	"net/http"
)

//Register handles user registration
type Register struct {
	db *sql.DB
}

//Register handles register route
func (reg *Register) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	utils.CheckParse(err, w, "Bad request", http.StatusBadRequest)

}

func CrawlURLs(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		err := req.ParseForm()
		if err != nil {
			http.Error(w, "Couldnt parse form", 502)
		}
		//url := req.FormValue("url")

	} else {
		http.Redirect(w, req, "/", http.StatusMethodNotAllowed)
	}
}
