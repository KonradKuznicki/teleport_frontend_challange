package main

import (
	"challenge/auth"
	"challenge/files"
	"challenge/server"
	"log"
	"net/http"
)

func main() {

	log.Println("starting")

	mux := SetupRouter()

	server.ServeTLS(mux, "../resources/server.ecdsa.crt", "../resources/server.ecdsa.key")
}

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	authenticator := auth.NewAuth()
	server.Handle(mux, "/user/logout", authenticator.LogoutHandler)
	server.Handle(mux, "/user/login", authenticator.LoginHandler)
	server.Handle(mux, "/login", auth.StaticsHandler)
	server.Handle(mux, "/files", auth.Wrapper(files.SataticsHandler))
	server.Handle(mux, "/", IndexHandler)
	return mux
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/files", http.StatusPermanentRedirect)
	} else {
		http.NotFound(w, r)
	}
}
