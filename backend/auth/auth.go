package auth

import (
	"net/http"
)

type Auth struct {
}

func (a *Auth) LoginHandler(writer http.ResponseWriter, request *http.Request) {
	http.SetCookie(writer, &http.Cookie{
		Name:   "auth",
		Value:  "true",
		MaxAge: 300,
		Path:   "/",
	})
	http.Redirect(writer, request, "/files", http.StatusFound)
}

func NewAuth() *Auth {
	return &Auth{}
}
