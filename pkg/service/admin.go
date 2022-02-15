package service

import (
	"database/sql"
	"fmt"
	"html/template"
	"knocker/pkg/repository"
	"knocker/pkg/sessionmanager"
	"knocker/pkg/tools"
	"net/http"
	"strings"
)

func Admin(w http.ResponseWriter, r *http.Request) {

	if !sessionmanager.Check_admcookie(r) {
		http.Redirect(w, r, "/admin/sign", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("templates/admin/admin.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	tmpl.ExecuteTemplate(w, "admin", nil)
}
func AdminSign(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/admin/adminsign.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl.ExecuteTemplate(w, "adminsign", nil)
}

func AdminLogin(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")
	if ok, err := repository.Admin_AUTH(login, password); ok {
		if !sessionmanager.Set_admcookie(w, r) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
		}
	} else {
		if err == sql.ErrNoRows {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		} else {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}
	}
}

func Admin_Get_Merch_List(w http.ResponseWriter, r *http.Request) {
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

func Admin_Get_Prod_List(w http.ResponseWriter, r *http.Request) {
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

func Admin_Get_Users_List(w http.ResponseWriter, r *http.Request) {
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

func Admin_Get_Orders_List(w http.ResponseWriter, r *http.Request) {
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
func Admin_Get_Admins_List(w http.ResponseWriter, r *http.Request) {
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
