package sessionmanager

import (
	"knocker/pkg/tools"
	"net/http"

	"github.com/gorilla/sessions"
)

func get_session(r *http.Request) *sessions.Session {
	var sessionsStore = sessions.NewCookieStore([]byte("secret"))
	session, err := sessionsStore.Get(r, "session")
	if err != nil {
		tools.Logger.Error(err.Error())
	}
	return session
}
