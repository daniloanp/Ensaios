package server

import (
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"

)



func handleStaticServer (mux *mux.Router) {
	//handler := http.StripPrefix("/static/", ))
	handlePreffix(mux, "/static/", http.FileServer(http.Dir("./frontend/build/web/static")))
	//mux.Handle(	"/static/", handler)
}

func handlePreffix(mux *mux.Router, url string, handler http.Handler ) {
	mux.PathPrefix(url).Handler(http.StripPrefix(url, handler))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
		bytes, err := ioutil.ReadFile("./frontend/build/web/index.html")
		if (err != nil) {
			w.Write([]byte("Erro!"))
		} else {
			w.Write(bytes)
		}
}


func BackendHandler() http.Handler {
	mux := mux.NewRouter()
	handleStaticServer(mux)
	handlePreffix(mux, "/", http.HandlerFunc(handleIndex))

	return mux
}

