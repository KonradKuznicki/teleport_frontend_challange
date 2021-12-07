package auth

import (
	"errors"
	"log"
	"net/http"
	"path"
	"strings"
)

func GoToLogin(writer http.ResponseWriter, request *http.Request) {
	http.SetCookie(writer, &http.Cookie{
		Name:     "returnTo",
		Value:    ComputeReturn(request.URL.Path),
		MaxAge:   300,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
	http.Redirect(writer, request, "/login", http.StatusFound)
}

func ComputeReturn(requestPath string) string {
	cleanPath := path.Clean(requestPath)
	if strings.Index(cleanPath, "/files") == 0 {
		return cleanPath
	}
	return "/files"
}

func (a *Auth) Wrapper(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		cookie, err := request.Cookie("auth")
		if err != nil && !errors.Is(err, http.ErrNoCookie) {
			log.Printf("error reading cookie: %v", err)
		}

		if cookie != nil {
			valid, err := a.IsTokenValid(cookie.Value)
			if err != nil {
				log.Println("error validating token ", err.Error())
				GoToLogin(writer, request)
				return
			}
			if valid {
				log.Println("valid user")
				handlerFunc(writer, request)
				return
			}
			log.Println("invalid token ")
			GoToLogin(writer, request)
			return
		} else {
			log.Println("redirecting")
			GoToLogin(writer, request)
		}
	}
}

func (a *Auth) WrapperAPI(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		cookie, err := request.Cookie("auth")
		if err != nil && !errors.Is(err, http.ErrNoCookie) {
			log.Printf("wierd! error reading cookie! %v", err)
		}

		if cookie != nil {
			valid, err := a.IsTokenValid(cookie.Value)
			if err != nil {
				log.Println("error validating token ", err.Error())
				http.Error(writer, "access denied", http.StatusForbidden)
				return
			}
			if valid {
				log.Println("valid user")
				handlerFunc(writer, request)
				return
			}
			log.Println("invalid token ")
			http.Error(writer, "access denied", http.StatusForbidden)
			return
		} else {
			log.Println("no cookie denied")
			http.Error(writer, "access denied", http.StatusForbidden)
		}
	}
}
