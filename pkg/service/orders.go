package service

import (
	"fmt"
	"html/template"
	"knocker/pkg/repository"
	"knocker/pkg/tools"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Admin_Get_Orders_List(w http.ResponseWriter, r *http.Request) {
	order_list, err := repository.Select_order_list()
	if tools.ErrorManager(err, w) {
		return
	}

	tmpl, err := template.ParseFiles("templates/admin/order_list.html", "templates/header.html", "templates/footer.html")
	if tools.ErrorManager(err, w) {
		return
	}
	tmpl.ExecuteTemplate(w, "order_list", order_list)
}

func Admin_Order_Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	order_id, err := strconv.ParseInt(vars["order_id"], 10, 32)
	tools.Logger.Trace("order_id: ", order_id)
	if err != nil {
		tools.Logger.Error(err.Error())
	}

	if repository.Delete_Order(int(order_id)) != nil {
		tools.Logger.Error(err.Error())
		return
	}
	href := fmt.Sprint("/admin/orders")
	http.Redirect(w, r, href, http.StatusMovedPermanently)
}
func Order_Statusup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	order_id, err := strconv.ParseInt(vars["order_id"], 10, 32)
	if err != nil {
		tools.Logger.Error(err.Error())
	}
	tools.Logger.Trace("order_id: ", order_id)

	if err := repository.Status_Order_up(int(order_id)); err != nil {
		tools.Logger.Error(err.Error())
		return
	}
	href := fmt.Sprintf("/admin/orders")
	tools.Logger.Tracef("redirect ", href)
	// http.Redirect(w, r, href, http.StatusMovedPermanently)
}
