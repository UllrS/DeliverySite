package sessionmanager

import (
	"knocker/pkg/tools"
	"net/http"

	"github.com/gorilla/sessions"
)

func get_session(r *http.Request) *sessions.Session {
	tools.Logger.Trace("start function")
	var sessionsStore = sessions.NewCookieStore([]byte("secret"))
	session, err := sessionsStore.Get(r, "session")
	tools.Logger.Trace(session)
	tools.Logger.Trace(err)
	if err != nil {
		tools.Logger.Error(err.Error())
	}
	tools.Logger.Trace("end function")
	return session
}
