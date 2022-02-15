package sessionmanager

import (
	"net/http"

	"github.com/gorilla/sessions"
)

func get_session(r *http.Request) *sessions.Session {
	var sessionsStore = sessions.NewCookieStore([]byte("secret"))
	session, _ := sessionsStore.Get(r, "session")
	return session
}
