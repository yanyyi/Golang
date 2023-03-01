package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	User    string
	Threads []string
}

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
			<head><title>Go Web Programming</title></head>
			<body><h1>Hello World</h1></body>
			</html>`
	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service, try next door")
}

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://www.baidu.com")
	w.WriteHeader(302)
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "Sau Sheong",
		Threads: []string{"fitst", "second", "third"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}
