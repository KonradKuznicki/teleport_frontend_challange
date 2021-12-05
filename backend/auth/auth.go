package auth

import (
	"net/http"
	"time"
)

type Auth struct {
}

func (a *Auth) LoginHandler(writer http.ResponseWriter, request *http.Request) {
	http.SetCookie(writer, &http.Cookie{
		Name:     "auth",
		Value:    "true",
		MaxAge:   300,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
	http.Redirect(writer, request, "/files", http.StatusFound)
}

func (a *Auth) LogoutHandler(writer http.ResponseWriter, request *http.Request) {
	http.SetCookie(writer, &http.Cookie{
		Name:     "auth",
		Value:    "false",
		Expires:  time.Unix(0, 0),
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
	http.Redirect(writer, request, "/login", http.StatusFound)
}

func NewAuth() *Auth {
	return &Auth{}
}
