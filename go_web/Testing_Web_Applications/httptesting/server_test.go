package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGet(t *testing.T) {
	if err := initSQL(); err != nil {
		fmt.Printf("connect to db failed, err:%v\n", err)
	}
	defer db.Close()
	mux := http.NewServeMux()
	mux.HandleFunc("/get/", handleGet)
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/get/1", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 1 {
		t.Errorf("Cannot retrieve JSON post")
	}
}
