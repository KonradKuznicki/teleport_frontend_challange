package main

import (
	"challenge/auth"
	"challenge/files"
	"log"
	"net/http"
)

func main() {
	// Set routing rules

	authenticator := auth.NewAuth()

	http.HandleFunc("/user/login", authenticator.LoginHandler)
	http.HandleFunc("/login", auth.StaticsHandler)
	http.HandleFunc("/files", auth.Wrapper(files.SataticsHandler))
	http.HandleFunc("/", IndexHandler)

	//Use the default DefaultServeMux.
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/files", http.StatusPermanentRedirect)
	} else {
		http.NotFound(w, r)
	}
}
