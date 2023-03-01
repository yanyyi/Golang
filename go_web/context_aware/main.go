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
	t, _ := template.ParseFiles("context_aware/tmpl.html")
	content := `I asked: <i>"What'' up?"</i>`
	t.Execute(w, content)
}
