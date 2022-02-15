package tools

import (
	"log"
	"net/http"
)

func ErrorManager(err error, w http.ResponseWriter) bool {
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return true
	} else {
		return false
	}

}
