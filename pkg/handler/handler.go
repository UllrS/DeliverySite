package handler

import (
	"knocker/pkg/service"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/admin", service.Admin).Methods("GET")
	rtr.HandleFunc("/admin/sign", service.AdminSign).Methods("GET")
	rtr.HandleFunc("/admin/login", service.AdminLogin).Methods("POST")
	rtr.HandleFunc("/admin/merchants", service.Admin_Get_Merch_List).Methods("GET")
	rtr.HandleFunc("/admin/products", service.Admin_Get_Prod_List).Methods("GET")
	rtr.HandleFunc("/admin/users", service.Admin_Get_Users_List).Methods("GET")
	rtr.HandleFunc("/admin/admins", service.Admin_Get_Admins_List).Methods("GET")
	rtr.HandleFunc("/admin/orders", service.Admin_Get_Orders_List).Methods("GET")

	rtr.HandleFunc("/create_product", service.Create_prod).Methods("GET")
	rtr.HandleFunc("/create_merchant", service.Create_merch).Methods("GET")
	rtr.HandleFunc("/save_product", service.Save_prod).Methods("POST")
	rtr.HandleFunc("/save_merchant", service.Save_merch).Methods("POST")

	rtr.HandleFunc("/admin/merchant/{merch_id:[0-9]+}", service.Detail_merchant).Methods("GET")

	rtr.HandleFunc("/", service.Index).Methods("GET")
	rtr.HandleFunc("/index/", service.Index).Methods("GET")
	rtr.HandleFunc("/about", service.About).Methods("GET")

	rtr.HandleFunc("/allmerchants", service.Get_Merch_List).Methods("GET")
	rtr.HandleFunc("/public_catering", service.Get_Merch_List).Methods("GET")
	rtr.HandleFunc("/grocery_stores", service.Get_Merch_List).Methods("GET")
	rtr.HandleFunc("/pharmacies", service.Get_Merch_List).Methods("GET")
	rtr.HandleFunc("/household_stores", service.Get_Merch_List).Methods("GET")
	rtr.HandleFunc("/transfer", service.Get_Merch_List).Methods("GET")
	rtr.HandleFunc("/merchant/{merch_id:[0-9]+}", service.Detail_merchant).Methods("GET")
	rtr.HandleFunc("/merchant/{merch_id:[0-9]+}/product/{prod_id:[0-9]+}", service.Detail_product).Methods("GET")
	rtr.HandleFunc("/basket", service.Basket).Methods("GET")
	rtr.HandleFunc("/basket_add", service.Basket_add).Methods("POST")
	rtr.HandleFunc("/Basket_delete", service.Basket_delete).Methods("POST")
	rtr.HandleFunc("/checkout", service.Checkout).Methods("POST")

	http.Handle("/", rtr)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.ListenAndServe(":5000", nil)
}
