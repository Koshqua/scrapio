package utils

import "net/http"

func CheckParse(err error, w http.ResponseWriter, errStr string, code int) {
	if err != nil {
		http.Error(w, errStr, code)
	}
}
