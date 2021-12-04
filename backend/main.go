package main

import (
	"challenge/auth"
	"challenge/files"
	"log"
	"net/http"
)

func main() {
	// Set routing rules
	http.HandleFunc("/login", auth.StaticsHandler)
	http.HandleFunc("/files", auth.Wrapper(files.SataticsHandler))
	http.HandleFunc("/", Tmp)

	//Use the default DefaultServeMux.
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Tmp(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/files", http.StatusPermanentRedirect)
	} else {
		http.NotFound(w, r)
	}
}
