package session

import (
	"github.com/gorilla/sessions"
	"github.com/daniloanp/Ensaios/application/backend/model"
	"net/http"
	"encoding/gob"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))


const (
	sessionName = "Session_Name"
	sessionErrorMessage = "Couldn't create session!!!"
)
type Session struct {
	*sessions.Session
	User model.UserAccount // maybe not
	Role model.RoleData // maybe not
}

func GetSessionData(w http.ResponseWriter,   r *http.Request) *Session {
	session, err := store.Get(r, sessionName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}
	if session.IsNew {
		session.Values["data"] = &Session{
			Session.Session: session,
			User: nil,
			Role: model.DbMap.Role
		}
	}
	return nil
}

func init() {
	gob.Register(&Session{})
}