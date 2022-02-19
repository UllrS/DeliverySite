package sessionmanager

import (
	"knocker/pkg/tools"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tools.Logger.Info(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func AdminAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if Check_admcookie(r) {
			next.ServeHTTP(w, r)
		} else {
			var uri = "/admin/sign"
			if r.RequestURI != uri {
				http.Redirect(w, r, uri, http.StatusSeeOther)
			}
			//http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	})
}
