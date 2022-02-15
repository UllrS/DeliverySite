package service

import (
	"fmt"
	"html/template"
	"knocker/pkg/repository"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func Get_Merch_List(w http.ResponseWriter, r *http.Request) {
	var field string
	var field_value string
	field = "type"
	field_value = strings.ReplaceAll(r.RequestURI, "/", "")
	if field_value == "allmerchants" {
		field_value = "%"
	}

	merch_list, err := repository.Select_merch_list_where(field, field_value)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	tmpl, err := template.ParseFiles("templates/merchant_list.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl.ExecuteTemplate(w, "merchant_list", merch_list)
}
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

func Detail_product(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	product_detail, err := repository.Select_product_where("id", vars["prod_id"])
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl, err := template.ParseFiles("templates/details_product.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	tmpl.ExecuteTemplate(w, "details", product_detail)

}
func Detail_merchant(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	product_list, err := repository.Select_product_list_where("merch", vars["merch_id"])
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl, err := template.ParseFiles("templates/product_list.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	tmpl.ExecuteTemplate(w, "product_list", product_list)

}
