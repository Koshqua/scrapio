package server

import (
	"database/sql"
	"github.com/koshqua/scrapio/utils"
	"net/http"
)

type Register struct {
	db *sql.DB
}

//Register handles register route
func (reg *Register) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	utils.CheckParse(err, w, "Bad request", http.StatusBadRequest)

}
