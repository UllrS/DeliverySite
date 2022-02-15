package service

import (
	"fmt"
	"html/template"
	"knocker/models"
	"knocker/pkg/repository"
	"knocker/pkg/sessionmanager"
	"knocker/pkg/tools"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
)

func Basket(w http.ResponseWriter, r *http.Request) {

	var sessionsStore = sessions.NewCookieStore([]byte("secret"))
	session, _ := sessionsStore.Get(r, "session")
	var basket_map = tools.JsonToMap(session.Values["basket"].([]byte))

	product_list, sum, err := repository.Basket_load("id", basket_map)
	fmt.Println("product_list")
	fmt.Println(product_list)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	tmpl, err := template.ParseFiles("templates/basket.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Println(sum)

	tmpl.ExecuteTemplate(w, "basket", product_list)
}
func Basket_add(w http.ResponseWriter, r *http.Request) {
	prod_id, err := strconv.ParseInt(r.FormValue("prod_id"), 10, 32)
	err = sessionmanager.Add_update_basket_cookie(w, r, prod_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func Basket_delete(w http.ResponseWriter, r *http.Request) {
	prod_id, _ := strconv.ParseInt(r.FormValue("prod_id"), 10, 32)

	err := sessionmanager.Add_update_basket_cookie(w, r, prod_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Checkout(w http.ResponseWriter, r *http.Request) {
	var user int
	tel := r.FormValue("tel")
	billing := r.FormValue("billing")

	var sessionsStore = sessions.NewCookieStore([]byte("secret"))
	session, _ := sessionsStore.Get(r, "session")
	var basket_map = tools.JsonToMap(session.Values["basket"].([]byte))

	product_list, sum, err := repository.Basket_load("id", basket_map)
	fmt.Println("product_list")
	fmt.Println(product_list)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	var order_list string
	order_list = fmt.Sprint("Summary: ", sum, "\n")
	order_list = fmt.Sprint(order_list, "Telephone: ", tel, "\n")
	order_list = fmt.Sprint(order_list, "Billing: ", billing, "\n")
	for i, prod := range product_list {
		order_list = fmt.Sprint(order_list, (i + 1), prod.Name, prod.Price, "\n")
	}
	user = sessionmanager.Get_User(w, r)
	fmt.Println("user")
	fmt.Println(user)
	var order_struct = models.Order{
		User:    user,
		Tel:     tel,
		Billing: billing,
		Basket:  order_list,
	}
	repository.Insert_order(order_struct)
	tmpl, err := template.ParseFiles("templates/checkout.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Println(sum)

	tmpl.ExecuteTemplate(w, "checkout", product_list)
}
