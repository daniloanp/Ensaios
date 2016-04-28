package main

import (

	"net/http"
	"io/ioutil"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/static/",http.StripPrefix("/static/", http.FileServer(http.Dir("./frontend/build/web/static"))))
	mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		bytes, err := ioutil.ReadFile("./frontend/build/web/index.html")
		if (err != nil) {
			w.Write([]byte("Erro!"))
		} else {
			w.Write(bytes)
		}
	})
	http.ListenAndServe(":8086", mux)
}