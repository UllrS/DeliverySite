package service

import (
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl.ExecuteTemplate(w, "index", nil)
}
func About(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/about.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl.ExecuteTemplate(w, "about", nil)
}
