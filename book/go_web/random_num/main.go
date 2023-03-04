package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/process", process)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("random_num/tmpl.html")
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)
}
