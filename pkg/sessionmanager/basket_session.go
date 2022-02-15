package sessionmanager

import (
	"encoding/json"
	"knocker/pkg/repository"
	"knocker/pkg/tools"
	"net/http"
)

func Add_update_basket_cookie(w http.ResponseWriter, r *http.Request, prod_id int64) error {
	session := get_session(r)

	var basket_map map[int]int
	if session.Values["basket"] == nil {
		basket_map = map[int]int{}
	} else {
		basket_map = tools.JsonToMap(session.Values["basket"].([]byte))
	}

	if val, found := basket_map[int(prod_id)]; found {
		basket_map[int(prod_id)] = val + 1
	} else {
		basket_map[int(prod_id)] = 1
	}

	j, _ := json.Marshal(basket_map)
	session.Values["basket"] = j
	// serialize and save cookie

	err := session.Save(r, w)
	if err != nil {
		return err
	}
	// save basket in database
	if session.Values["user_id"] != nil && session.Values["user_id"] != 0 {
		var i = session.Values["user_id"].(int)
		err = repository.Basket_update(i, string(j))
	}
	return err
}
func Del_update_basket_cookie(w http.ResponseWriter, r *http.Request, prod_id int64) error {
	session := get_session(r)

	var basket_map map[int]int
	if session.Values["basket"] == nil {
		basket_map = map[int]int{}
	} else {
		basket_map = tools.JsonToMap(session.Values["basket"].([]byte))
	}

	if val, ok := basket_map[int(prod_id)]; ok && val > 1 {
		basket_map[int(prod_id)] = val - 1
	} else {
		delete(basket_map, int(prod_id))
	}

	j, _ := json.Marshal(basket_map)
	session.Values["basket"] = j
	// serialize and save cookie

	err := session.Save(r, w)
	if err != nil {
		return err
	}
	// save basket in database
	if session.Values["user_id"] != nil && session.Values["user_id"] != 0 {
		var i = session.Values["user_id"].(int)
		err = repository.Basket_update(i, string(j))
	}
	return err
}
func Get_User(w http.ResponseWriter, r *http.Request) int {
	session := get_session(r)
	var user int
	user = 0
	if session.Values["user_id"] != nil && session.Values["user_id"] != 0 {
		user = session.Values["user_id"].(int)
		return user
	}
	return user
}
