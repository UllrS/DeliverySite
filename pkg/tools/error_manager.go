package tools

import (
	"net/http"
)

func ErrorManager(err error, w http.ResponseWriter) bool {
	if err != nil {
		Logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return true
	} else {
		return false
	}

}
