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
	"time"
)

func Basket(w http.ResponseWriter, r *http.Request) {

	basket_map := sessionmanager.Get_Basket_Map(w, r)

	product_list, sum, err := repository.Basket_load("id", basket_map)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		tools.Logger.Error(err.Error())
		return
	}

	tmpl, err := template.ParseFiles("templates/basket.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		tools.Logger.Error(err.Error())
		return
	}
	tools.Logger.Tracef("Basket summary: ", sum)

	tmpl.ExecuteTemplate(w, "basket", product_list)
}
func Basket_add(w http.ResponseWriter, r *http.Request) {
	prod_id, err := strconv.ParseInt(r.FormValue("prod_id"), 10, 32)
	//merch, err := strconv.ParseInt(r.FormValue("merch"), 10, 32)
	err = sessionmanager.Add_update_basket_cookie(w, r, prod_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "<script>alert('Item added to cart'); window.location.replace('/basket');</script>")
	// http.Redirect(w, r, "/basket", http.StatusSeeOther)
	// w.Write([]byte("<script>alert('Item added')</script>"))
}
func Basket_delete(w http.ResponseWriter, r *http.Request) {
	prod_id, _ := strconv.ParseInt(r.FormValue("prod_id"), 10, 32)
	//merch, _ := strconv.ParseInt(r.FormValue("merch"), 10, 32)

	err := sessionmanager.Del_update_basket_cookie(w, r, prod_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		tools.Logger.Error(err.Error())
		return
	}

	http.Redirect(w, r, "/basket", http.StatusSeeOther)
}

func Checkout(w http.ResponseWriter, r *http.Request) {
	var user int
	tel := r.FormValue("tel")
	shipping := r.FormValue("shipping")
	date := time.Now().Format("2006-01-02")
	if !(tools.TelValidator(tel) && tools.BilingValidator(shipping)) {
		http.Error(w, "Enter telephone and delivery address", http.StatusInternalServerError)
		return

	}
	basket_map := sessionmanager.Get_Basket_Map(w, r)
	if len(basket_map) < 1 {
		http.Error(w, "Add item to cart", http.StatusInternalServerError)
		return
	}

	product_list, sum, err := repository.Basket_load("id", basket_map)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		tools.Logger.Error(err.Error())
		return
	}
	var order_list string
	order_list = fmt.Sprint("Summary: ", sum, "\n")
	order_list = fmt.Sprint(order_list, "Telephone: ", tel, "\n")
	order_list = fmt.Sprint(order_list, "Shipping: ", shipping, "\n")
	for i, prod := range product_list {
		order_list = fmt.Sprint(order_list, (i + 1), prod.Name, prod.Price, "\n")
	}
	user = sessionmanager.Get_User(w, r)
	var order_struct = models.Order{
		User:     user,
		Tel:      tel,
		Shipping: shipping,
		Basket:   order_list,
		Date:     date,
	}
	if tools.ErrorManager(repository.Insert_order(order_struct), w) {
		return
	}
	sessionmanager.Clear_Basket_Map(w, r)

	tmpl, err := template.ParseFiles("templates/checkout.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		tools.Logger.Error(err.Error())
		return
	}
	tools.Logger.Tracef("Basket summary: ", sum)

	tmpl.ExecuteTemplate(w, "checkout", product_list)
}
