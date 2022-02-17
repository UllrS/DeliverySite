package service

import (
	"database/sql"
	"fmt"
	"html/template"
	"knocker/pkg/repository"
	"knocker/pkg/sessionmanager"
	"knocker/pkg/tools"
	"net/http"
)

func Admin(w http.ResponseWriter, r *http.Request) {

	tools.Logger.Trace("start function")

	tmpl, err := template.ParseFiles("templates/admin/admin.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	tmpl.ExecuteTemplate(w, "admin", nil)
	tools.Logger.Trace("end function")
}
func AdminSign(w http.ResponseWriter, r *http.Request) {
	tools.Logger.Trace("start function")
	tmpl, err := template.ParseFiles("templates/admin/adminsign.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl.ExecuteTemplate(w, "adminsign", nil)
	tools.Logger.Trace("end function")
}

func AdminLogin(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")
	if ok, err := repository.Admin_AUTH(login, password); ok {
		if !sessionmanager.Set_admcookie(w, r) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			tools.Logger.Warn(err.Error())
		} else {
			http.Redirect(w, r, "/admin/", http.StatusSeeOther)
		}
	} else {
		if err == sql.ErrNoRows {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			tools.Logger.Error("ADMIN BAD PWD")
		} else {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			tools.Logger.Error(err.Error())
		}
	}
}
