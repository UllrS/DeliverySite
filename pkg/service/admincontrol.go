package service

import (
	"bufio"
	"html/template"
	"knocker/pkg/repository"
	"knocker/pkg/tools"
	"net/http"
	"os"
)

func Admin_list(w http.ResponseWriter, r *http.Request) {
	admin_list, err := repository.Select_admin_list()
	if tools.ErrorManager(err, w) {
		return
	}

	tmpl, err := template.ParseFiles("templates/admin/admin_list.html", "templates/header.html", "templates/footer.html")
	if tools.ErrorManager(err, w) {
		return
	}
	tmpl.ExecuteTemplate(w, "admin_list", admin_list)
}

func User_list(w http.ResponseWriter, r *http.Request) {
	user_list, err := repository.Select_User_List()
	if tools.ErrorManager(err, w) {
		return
	}

	tmpl, err := template.ParseFiles("templates/admin/user_list.html", "templates/header.html", "templates/footer.html")
	if tools.ErrorManager(err, w) {
		return
	}
	tmpl.ExecuteTemplate(w, "user_list", user_list)
}
func Admin_Logs(w http.ResponseWriter, r *http.Request) {
	// user_list, err := repository.Select_User_List()
	// if tools.ErrorManager(err, w) {
	// 	return
	// }
	var log_list []string
	file, err := os.Open("logs/list.log")
	if err != nil {
		tools.Logger.Error(err.Error())
	}
	// wr := bytes.Buffer{}
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		log_list = append(log_list, sc.Text())
	}
	var log_list_reverse = tools.ReverseArray(log_list)
	defer file.Close()
	tmpl, err := template.ParseFiles("templates/admin/logs.html", "templates/header.html", "templates/footer.html")
	if tools.ErrorManager(err, w) {
		return
	}
	tmpl.ExecuteTemplate(w, "logs", log_list_reverse)
}
