package sessionmanager

import (
	"net/http"
)

func Check_admcookie(r *http.Request) bool {
	return get_session(r).Values["adm"] == 1
}
func Set_admcookie(w http.ResponseWriter, r *http.Request) bool {
	session := get_session(r)
	session.Values["adm"] = 1
	err := session.Save(r, w)
	if err == nil {
		return true
	}
	return false
}
