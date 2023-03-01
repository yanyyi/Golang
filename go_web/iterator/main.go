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
	t, _ := template.ParseFiles("go_web/iterator/tmpl.html")
	seasons := []string{"Spring", "Summer", "Autumn", "Winter"}
	t.Execute(w, seasons)
}
