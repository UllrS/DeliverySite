package handler

import (
	"knocker/pkg/service"
	"knocker/pkg/sessionmanager"
	"knocker/pkg/tools"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	rtr := mux.NewRouter()
	auth_subrouter := rtr.PathPrefix("/admin").Subrouter()
	auth_subrouter.Use(sessionmanager.LoggingMiddleware)
	auth_subrouter.Use(sessionmanager.AuthMiddleware)

	auth_subrouter.HandleFunc("/", service.Admin).Methods("GET")

	auth_subrouter.HandleFunc("/logs", service.Admin_Logs).Methods("GET")
	auth_subrouter.HandleFunc("/merchants", service.Adm_Merch_List).Methods("GET")
	auth_subrouter.HandleFunc("/users", service.User_list).Methods("GET")
	auth_subrouter.HandleFunc("/admins", service.Admin_list).Methods("GET")
	auth_subrouter.HandleFunc("/orders", service.Admin_Get_Orders_List).Methods("GET")

	auth_subrouter.HandleFunc("/merchant/create", service.Create_merch).Methods("GET")
	auth_subrouter.HandleFunc("/merchant/save", service.Save_merch).Methods("POST")
	auth_subrouter.HandleFunc("/merchant/{merch_id:[0-9]+}", service.Admin_Prod_List).Methods("GET")
	auth_subrouter.HandleFunc("/merchant/{merch_id:[0-9]+}/edit", service.Admin_Merch_Edit).Methods("GET")
	auth_subrouter.HandleFunc("/merchant/{merch_id:[0-9]+}/edit", service.Edit_merch).Methods("POST")
	auth_subrouter.HandleFunc("/merchant/{merch_id:[0-9]+}/delete", service.Admin_Merch_Delete).Methods("GET")

	auth_subrouter.HandleFunc("/merchant/{merch_id:[0-9]+}/product/create", service.Create_prod).Methods("GET")
	auth_subrouter.HandleFunc("/merchant/{merch_id:[0-9]+}/product/save", service.Save_prod).Methods("POST")
	auth_subrouter.HandleFunc("/merchant/{merch_id:[0-9]+}/product/{prod_id:[0-9]+}", service.Detail_product).Methods("GET")
	auth_subrouter.HandleFunc("/merchant/{merch_id:[0-9]+}/product/{prod_id:[0-9]+}/edit", service.Admin_Prod_Edit).Methods("GET")
	auth_subrouter.HandleFunc("/merchant/{merch_id:[0-9]+}/product/{prod_id:[0-9]+}/delete", service.Admin_Prod_Delete).Methods("GET")
	auth_subrouter.HandleFunc("/merchant/{merch_id:[0-9]+}/product/{prod_id:[0-9]+}/edit", service.Edit_Prod).Methods("POST")

	auth_subrouter.HandleFunc("/order/delete/{order_id:[0-9]+}", service.Admin_Order_Delete).Methods("GET")
	auth_subrouter.HandleFunc("/order/status/{order_id:[0-9]+}", service.Order_Statusup).Methods("GET")

	rtr.HandleFunc("/admin/sign", service.AdminSign).Methods("GET")
	rtr.HandleFunc("/admin/login", service.AdminLogin).Methods("POST")

	rtr.HandleFunc("/", service.Index).Methods("GET")
	rtr.HandleFunc("/index/", service.Index).Methods("GET")
	rtr.HandleFunc("/about", service.About).Methods("GET")

	rtr.HandleFunc("/merchant/allmerchants", service.Get_Merch_List).Methods("GET")
	rtr.HandleFunc("/merchant/public_catering", service.Get_Merch_List).Methods("GET")
	rtr.HandleFunc("/merchant/grocery_stores", service.Get_Merch_List).Methods("GET")
	rtr.HandleFunc("/merchant/pharmacies", service.Get_Merch_List).Methods("GET")
	rtr.HandleFunc("/merchant/household_stores", service.Get_Merch_List).Methods("GET")
	rtr.HandleFunc("/merchant/transfer", service.Get_Merch_List).Methods("GET")
	rtr.HandleFunc("/merchant/{merch_id:[0-9]+}", service.Detail_merchant).Methods("GET")
	rtr.HandleFunc("/merchant/{merch_id:[0-9]+}/product/{prod_id:[0-9]+}", service.Detail_product).Methods("GET")

	rtr.HandleFunc("/basket", service.Basket).Methods("GET")
	rtr.HandleFunc("/basket/add", service.Basket_add).Methods("POST")
	rtr.HandleFunc("/basket/delete", service.Basket_delete).Methods("POST")
	rtr.HandleFunc("/checkout", service.Checkout).Methods("POST")

	http.Handle("/", rtr)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	tools.Logger.Debug("Start SERVER")
	http.ListenAndServe(":5000", nil)
}
