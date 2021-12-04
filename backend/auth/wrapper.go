package auth

import "net/http"

func Wrapper(http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, "/login", http.StatusFound)
	}
}
