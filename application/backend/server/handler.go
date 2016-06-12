package server

import (
	"net/http"
	"github.com/daniloanp/Ensaios/application/backend/session"
)

type handler func (w http.ResponseWriter, r *http.Request, s session.Session)

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s := session.Clone()
	defer s.Close()

	h(w, r, s.DB("example"))
}