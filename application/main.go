package application

import (
	"github.com/daniloanp/Ensaios/application/backend/server"
	"net/http"
)

func main() {
	mux := server.BackendHandler()
	http.ListenAndServe(":8080", mux)
}