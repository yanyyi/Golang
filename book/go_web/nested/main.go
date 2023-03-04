package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/process", process)
	http.ListenAndServe(":8080", nil)
}

func process(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	if rand.Intn(10) > 5 {
		t, _ = template.ParseFiles("go_web/nested/layout.html", "go_web/nested/red_hello.html")
	} else {
		t, _ = template.ParseFiles("go_web/nested/layout.html", "go_web/nested/blue_hello.html")
	}
	//t, _ = template.ParseFiles("go_web/nested/layout.html", "go_web/nested/blue_hello.html")
	t.ExecuteTemplate(w, "layout", "")
}
