package auth

import (
	"log"
	"net/http"
)

func Wrapper(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		cookie, err := request.Cookie("auth")
		if err != nil {
			log.Printf("wierd! error reading cookie! %v", err)
		}
		if cookie != nil && cookie.Value == "true" {
			log.Println("valid user")
			handlerFunc(writer, request)
		} else {
			log.Println("redirecting")
			http.Redirect(writer, request, "/login", http.StatusFound)
		}
	}
}
