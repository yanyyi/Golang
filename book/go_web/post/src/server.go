package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
)

type Post struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	if err := initSQL(); err != nil {
		fmt.Printf("connect to db failed, err:%v\n", err)
	}
	defer db.Close()
	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	http.HandleFunc("/get/", handleGet)
	http.HandleFunc("/post/", handlePost)
	http.HandleFunc("/put/", handlePut)
	http.HandleFunc("/delete/", handleDelete)
	server.ListenAndServe()
}

// Retrieve a post
// GET /post/1
func handleGet(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("r.URL", r.URL)
	//fmt.Println("r.URL.Path:", r.URL.Path)
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := SelectAPost(id)
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// Create a post
// POST /post/
func handlePost(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var post Post
	json.Unmarshal(body, &post)
	err := post.Create()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

// Update a post
// PUT /post/1
func handlePut(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := SelectAPost(id)
	if err != nil {
		return
	}
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &post)
	err = post.Update()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

// // Delete a post
// DELETE /post/1
func handleDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := SelectAPost(id)
	if err != nil {
		return
	}
	err = post.Delete()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}
