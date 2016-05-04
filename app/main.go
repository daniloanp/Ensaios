package main

import (
	"github.com/daniloanp/Ensaios/app/backend/server"
	"net/http"
)

func main() {
	mux := server.BackendHandler()
	http.ListenAndServe(":8080", mux)
}