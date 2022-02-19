package service

import (
	"fmt"
	"html/template"
	"knocker/models"
	"knocker/pkg/repository"
	"knocker/pkg/tools"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func Create_merch(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/admin/create_merch.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		tools.Logger.Error(err.Error())
	}
	tmpl.ExecuteTemplate(w, "create_merch", nil)
}
func Create_prod(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	merch_id, err := strconv.ParseInt(vars["merch_id"], 10, 32)
	tmpl, err := template.ParseFiles("templates/admin/create_prod.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		tools.Logger.Error(err.Error())
	}
	tmpl.ExecuteTemplate(w, "create_prod", merch_id)
}

func Save_prod(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("title")
	anons := r.FormValue("anons")
	unit := r.FormValue("unit")
	category := r.FormValue("category")
	price, _ := strconv.ParseFloat(r.FormValue("price"), 32)
	merch, _ := strconv.ParseInt(r.FormValue("merch"), 10, 32)
	portion, _ := strconv.ParseFloat(r.FormValue("portion"), 32)
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
	tools.Logger.Tracef("prod_struct: ", prod_struct)
	var id int
	if id = repository.Insert_product(prod_struct); id == 0 {
		return
	}
	tools.LoadFile(w, r, id, "prod")

	http.Redirect(w, r, "/admin/", http.StatusSeeOther)
}
func Save_merch(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("title")
	merch_type := r.FormValue("merch_type")
	addr := r.FormValue("addr")
	anons := r.FormValue("anons")
	date := time.Now().Format("2006-01-02")

	if name == "" || merch_type == "" {
		fmt.Fprintf(w, "Не хватает данных")
		return
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

	http.Redirect(w, r, "/admin/", http.StatusSeeOther)
}
func Edit_merch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	merch_id, _ := strconv.ParseInt(vars["merch_id"], 10, 32)
	name := r.FormValue("title")
	merch_type := r.FormValue("merch_type")
	addr := r.FormValue("addr")
	anons := r.FormValue("anons")
	date := time.Now().Format("2006-01-02")

	if name == "" || merch_type == "" {
		fmt.Fprintf(w, "Не хватает данных")
		return
	}
	var merch_struct = models.Merchant{
		Id:    int32(merch_id),
		Name:  name,
		Type:  merch_type,
		Addr:  addr,
		Anons: anons,
		Img:   "",
		Date:  date,
	}

	if err := repository.Update_merchant(merch_struct); err != nil {
		return
	}
	tools.LoadFile(w, r, int(merch_struct.Id), "merch")

	http.Redirect(w, r, "/admin/", http.StatusSeeOther)
}
func Edit_Prod(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.ParseInt(r.FormValue("prod_id"), 10, 32)
	name := r.FormValue("title")
	anons := r.FormValue("anons")
	unit := r.FormValue("unit")
	category := r.FormValue("category")
	price, _ := strconv.ParseFloat(r.FormValue("price"), 32)
	merch, _ := strconv.ParseInt(r.FormValue("merch"), 10, 32)
	portion, _ := strconv.ParseFloat(r.FormValue("portion"), 32)
	if name == "" {
		fmt.Fprintf(w, "Не хватает данных")
	}

	var prod_struct = models.Product{
		Id:       int32(id),
		Name:     name,
		Price:    float32(price),
		Category: category,
		Anons:    anons,
		Merch:    int32(merch),
		Portion:  float32(portion),
		Unit:     unit,
	}
	tools.Logger.Tracef("prod_struct: ", prod_struct)
	if id == 0 {
		tools.Logger.Error("Bad id")
		return
	}
	if err := repository.Update_Product(prod_struct); err != nil {
		tools.Logger.Error(err.Error())
		return
	}
	tools.LoadFile(w, r, int(prod_struct.Id), "prod")

	http.Redirect(w, r, "/admin/", http.StatusSeeOther)
}

func Adm_Merch_List(w http.ResponseWriter, r *http.Request) {
	var field string
	var field_value string
	field = "type"
	field_value = strings.ReplaceAll(r.RequestURI, "/admin/", "")
	if field_value == "merchants" {
		field_value = "%"
	}
	merch_list, err := repository.Select_merch_list_where(field, field_value)
	if tools.ErrorManager(err, w) {
		return
	}

	tmpl, err := template.ParseFiles("templates/admin/admin_merchant_list.html", "templates/header.html", "templates/footer.html")
	if tools.ErrorManager(err, w) {
		return
	}
	tmpl.ExecuteTemplate(w, "admin_merchant_list", merch_list)
}

func Admin_Merch_Edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	merch_id, err := strconv.ParseInt(vars["merch_id"], 10, 32)
	if err != nil {
		tools.Logger.Error(err.Error())
	}

	merch, err := repository.Select_merch_where(int(merch_id))
	if err != nil {
		tools.Logger.Error(err.Error())
		return
	}
	tmpl, err := template.ParseFiles("templates/admin/edit_merch.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl.ExecuteTemplate(w, "edit_merch", merch)
}
func Admin_Prod_Edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//merch_id := vars["merch_id"]
	prod_id := vars["prod_id"]

	merch, err := repository.Select_product_where("id", prod_id)
	if err != nil {
		tools.Logger.Error(err.Error())
		return
	}
	tmpl, err := template.ParseFiles("templates/admin/edit_prod.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl.ExecuteTemplate(w, "edit_prod", merch)

}
func Admin_Merch_Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	merch_id, err := strconv.ParseInt(vars["merch_id"], 10, 32)
	if err != nil {
		tools.Logger.Error(err.Error())
		return
	}

	if repository.Delete_merch(int(merch_id)) != nil {
		tools.Logger.Error(err.Error())
		return
	}
	http.Redirect(w, r, "/admin/merchants", http.StatusMovedPermanently)

}
func Admin_Prod_Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	prod_id, err := strconv.ParseInt(vars["prod_id"], 10, 32)
	merch_id, err := strconv.ParseInt(vars["merch_id"], 10, 32)
	if err != nil {
		tools.Logger.Error(err.Error())
	}

	if repository.Delete_prod(int(prod_id)) != nil {
		tools.Logger.Error(err.Error())
		return
	}
	href := fmt.Sprintf("/admin/merchant/%d", merch_id)
	http.Redirect(w, r, href, http.StatusMovedPermanently)

}
func Admin_Prod_List(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	product_list, err := repository.Select_product_list_where("merch", vars["merch_id"])
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl, err := template.ParseFiles("templates/admin/admin_product_list.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	tmpl.ExecuteTemplate(w, "admin_product_list", product_list)

}
