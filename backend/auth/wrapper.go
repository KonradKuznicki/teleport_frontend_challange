package auth

import (
	"log"
	"net/http"
)

func (a *Auth) Wrapper(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		cookie, err := request.Cookie("auth")
		if err != nil {
			log.Printf("wierd! error reading cookie! %v", err)
		}

		if cookie != nil {
			valid, err := a.IsTokenValid(cookie.Value)
			if err != nil {
				log.Println("error validating token ", err.Error())
				http.Redirect(writer, request, "/login", http.StatusFound)
				return
			}
			if valid {
				log.Println("valid user")
				handlerFunc(writer, request)
				return
			}
			log.Println("invalid token ")
			http.Redirect(writer, request, "/login", http.StatusFound)
			return
		} else {
			log.Println("redirecting")
			http.Redirect(writer, request, "/login", http.StatusFound)
		}
	}
}
