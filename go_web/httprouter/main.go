package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	mux := httprouter.New()
	mux.GET("/hello/:name", hello)

	//server := http.Server{
	//	Addr:    "127.0.0.1:8080",
	//	Handler: mux,
	//}
	//server.ListenAndServe()
	http.ListenAndServe(":8080", mux)
}

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", p.ByName("name"))
}
