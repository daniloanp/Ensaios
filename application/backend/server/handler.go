package server

import (
	"net/http"
	"github.com/daniloanp/Ensaios/application/backend/session"
	"github.com/daniloanp/Ensaios/application/backend/app"
)

type handler func (w http.ResponseWriter, r *http.Request, s *session.Session)

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s := session.GetSessionData(w, r)
	if err, ok := app.Db().Role.HasPermission(s.Role, r.URL.Path); err != nil  {
		///TODO:ErrorAction
		http.Error(w, "Some error Ocurred", http.StatusInternalServerError)
	} else if (!ok) {
		///TODO:NoPermissionAction
		http.Error(w, "No Permission", http.StatusForbidden)
	}
	h(w, r, s)
}