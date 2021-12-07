package main

import (
	"challenge/auth"
	"challenge/auth/userRepositories"
	"challenge/files"
	"challenge/server"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {

	log.Println("starting")

	mux := SetupRouter()

	server.ServeTLS(mux, "../resources/server.ecdsa.crt", "../resources/server.ecdsa.key")
}

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	notVerySecurePass := fmt.Sprintf("%x", strings.Repeat("A", 32)) // TODO: should be some random 32byte

	authenticator := auth.NewAuth(
		userRepositories.NewInMemoryUserRepository(),
		auth.NewHardHasher(),
		auth.NewSessionManager(auth.NewAESEncryptor(notVerySecurePass), "magic secret", time.Minute),
		70)

	err := authenticator.CreateUser("user1", "pass1")
	if err != nil {
		log.Fatal(err)
	}
	err = authenticator.CreateUser("user2", "pass2")
	if err != nil {
		log.Fatal(err)
	}

	fm, err := files.NewFileManager("../resources/traversable")
	if err != nil {
		log.Fatal("cannot open file manager: %v", err)
	}
	server.Handle(mux, "/API/v1/user/logout", enableCors(authenticator.LogoutHandler))
	server.Handle(mux, "/API/v1/user/login", enableCors(authenticator.LoginHandler))
	server.Handle(mux, "/API/v1/files/", authenticator.Wrapper(enableCors(fm.FilesHandler)))
	server.Handle(mux, "/login/", auth.StaticsHandler)
	server.Handle(mux, "/files/", authenticator.Wrapper(files.SataticsHandler))
	server.Handle(mux, "/", IndexHandler(authenticator))
	return mux
}

func IndexHandler(authenticator *auth.Auth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.Redirect(w, r, "/files", http.StatusPermanentRedirect)
		} else {
			log.Println("magic magic")
			http.NotFound(w, r)
		}
	}
}

func enableCors(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		if request.Method != "OPTIONS" {
			handlerFunc(writer, request)
		}
	}
}
