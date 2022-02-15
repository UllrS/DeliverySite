package service

import (
	"fmt"
	"html/template"
	"knocker/models"
	"knocker/pkg/repository"
	"knocker/pkg/tools"
	"net/http"
	"strconv"
	"time"
)

func Create_merch(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/admin/create_merch.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl.ExecuteTemplate(w, "create_merch", nil)
}
func Create_prod(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/admin/create_prod.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl.ExecuteTemplate(w, "create_prod", nil)
}

func Save_prod(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("title")
	anons := r.FormValue("anons")
	unit := r.FormValue("unit")
	category := r.FormValue("category")
	price, _ := strconv.ParseFloat(r.FormValue("price"), 32)
	merch, _ := strconv.ParseInt(r.FormValue("merch"), 10, 32)
	portion, _ := strconv.ParseFloat(r.FormValue("portion"), 32)
	fmt.Println(price)
	if name == "" {
		fmt.Fprintf(w, "Не хватает данных")
	}

	var prod_struct = models.Product{
		Name:     name,
		Price:    float32(price),
		Category: category,
		Anons:    anons,
		Merch:    int32(merch),
		Unit:     unit,
		Portion:  float32(portion),
	}
	fmt.Println(prod_struct)
	var id int
	if id = repository.Insert_product(prod_struct); id == 0 {
		return
	}
	tools.LoadFile(w, r, id, "prod")

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func Save_merch(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("title")
	merch_type := r.FormValue("merch_type")
	addr := r.FormValue("addr")
	anons := r.FormValue("anons")
	date := time.Now().Format("2006-01-02")

	if name == "" || merch_type == "" {
		fmt.Fprintf(w, "Не хватает данных")
	}
	var merch_struct = models.Merchant{
		Name:  name,
		Type:  merch_type,
		Addr:  addr,
		Anons: anons,
		Img:   "",
		Date:  date,
	}
	var id int
	if id = repository.Insert_merchant(merch_struct); id == 0 {
		return
	}
	tools.LoadFile(w, r, id, "merch")

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
