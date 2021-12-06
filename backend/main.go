package main

import (
	"challenge/auth"
	"challenge/files"
	"challenge/server"
	"log"
	"net/http"
	"strings"
)

func main() {

	log.Println("starting")

	mux := SetupRouter()

	server.ServeTLS(mux, "../resources/server.ecdsa.crt", "../resources/server.ecdsa.key")
}

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	authenticator := auth.NewAuth()
	server.Handle(mux, "/API/v1/user/logout", enableCors(authenticator.LogoutHandler))
	server.Handle(mux, "/API/v1/user/login", enableCors(authenticator.LoginHandler))
	server.Handle(mux, "/login", auth.StaticsHandler)
	server.Handle(mux, "/files", auth.Wrapper(files.SataticsHandler))
	server.Handle(mux, "/", IndexHandler)
	return mux
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/files", http.StatusPermanentRedirect)
	} else if strings.Index(r.URL.Path, "/API/v1/files") == 0 {
		log.Println("files matched")
		enableCors(files.FilesHandler)(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func enableCors(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		handlerFunc(writer, request)
	}
}
