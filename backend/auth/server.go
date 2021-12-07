package auth

import (
	"net/http"
)

func StaticsHandler(writer http.ResponseWriter, request *http.Request) {

	fs := http.FileServer(http.Dir("../frontend/login-build"))
	http.StripPrefix("/login", fs).ServeHTTP(writer, request)
	// 	writer.Header().Add("content-type", "text/html")
	// 	_, err := io.WriteString(writer, "<html><head></head><body><h1>login form</h1><a href='/user/login'>log me in</a></body></html>")
	// 	if err != nil {
	// 		log.Printf("error writeing response for %s error: %v", request.URL.Host, err)
	// 	}
}
