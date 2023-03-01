package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	http.HandleFunc("/set_message", setMessage)
	http.HandleFunc("/show_message", showMessage)
	server.ListenAndServe()
}

func setMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("黑暗森林法则:黑暗森林法则可简单理解为，一旦某个宇宙文明被发现，就必然遭到其他宇宙文明的打击。")
	c := http.Cookie{
		Name:  "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c)
}

func showMessage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "No message found")

		}
	} else {
		rc := http.Cookie{
			Name:    "flash",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
}
