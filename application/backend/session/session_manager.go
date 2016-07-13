package session

import (
	"github.com/gorilla/sessions"
	"net/http"
	//"encoding/gob"
	"github.com/daniloanp/Ensaios/application/backend/app"
	"github.com/daniloanp/Ensaios/application/backend/model/tables"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))


const (
	dataKey = string("data")
	sessionName = "Session_Name"
	sessionErrorMessage = "Couldn't create session!!!"
)
type Session struct {
	*sessions.Session
	User *tables.UserAccountData // maybe not
	Role *tables.RoleData // maybe not
}

func GetSessionData(w http.ResponseWriter,   r *http.Request) *Session {
	var mySession *Session
	session, err := store.Get(r, sessionName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) //TODO:RESOLVE_IT
		return nil
	}
	if session.IsNew {
		role, _ := app.Db().Role().GetByID(app.AnonymousRole); //TODO:Ignoring Error
		mySession = &Session{
			Session: session,
			User: nil,
			Role: role ,
		}
		session.Values[dataKey] = mySession
		session.Save(r,w)
	} else {
		var (
			ok bool
			value = session.Values[dataKey]
		)
		if mySession, ok = value.(*Session); !ok {
			http.Error(w, err.Error(), http.StatusInternalServerError) //TODO:RESOLVE_IT
			return nil
		}
	}
	return mySession
}

func init() {
	//gob.Register(&Session{})
}