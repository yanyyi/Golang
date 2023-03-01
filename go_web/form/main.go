package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

func process(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	//r.ParseMultipartForm(1024)
	fmt.Fprintln(w, r.FormValue("hello"))
	fmt.Fprintln(w, r.Form)
}
