package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/process", process)
	http.ListenAndServe(":8080", nil)
}

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("go_web/include/t1.html", "go_web/include/t2.html")
	t.Execute(w, "Go Web编程教学")
}
