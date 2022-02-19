package sessionmanager

import (
	"knocker/pkg/tools"
	"net/http"
)

func Check_admcookie(r *http.Request) bool {
	tools.Logger.Warn("Admin cookie no found")
	return get_session(r).Values["adm"] == 1
}
func Set_admcookie(w http.ResponseWriter, r *http.Request) bool {
	session := get_session(r)
	session.Values["adm"] = 1
	err := session.Save(r, w)
	if err == nil {
		tools.Logger.Tracef("end function", true)
		return true
	}
	tools.Logger.Error(err.Error())
	return false
}
