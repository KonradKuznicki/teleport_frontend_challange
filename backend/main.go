package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// Set routing rules
	http.HandleFunc("/", Tmp)

	//Use the default DefaultServeMux.
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Tmp(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/files", 301)
	}
	if r.URL.Path == "/files" {
		cookie, err := r.Cookie("auth")
		if err != nil {
			log.Printf("cookie reading error?! %v", err)
		}

		if cookie == nil {
			http.Redirect(w, r, "/login", 301)
		} else if cookie.Value == "true" {
			w.Header().Add("content-type", "text/html")
			_, err := io.WriteString(w, "<html><head></head><body><h1>file manager</h1></body></html>")
			if err != nil {
				log.Printf("error writeing response for %s error: %v", r.URL.Host, err)
			}
		} else {

		}
	}
	if r.URL.Path == "/login" {
		w.Header().Add("content-type", "text/html")
		_, err := io.WriteString(w, "<html><head></head><body><h1>login form</h1></body></html>")
		if err != nil {
			log.Printf("error writeing response for %s error: %v", r.URL.Host, err)
		}
	}

}
