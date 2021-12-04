package auth

import (
	"io"
	"log"
	"net/http"
)

func StaticsHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("content-type", "text/html")
	_, err := io.WriteString(writer, "<html><head></head><body><h1>login form</h1></body></html>")
	if err != nil {
		log.Printf("error writeing response for %s error: %v", request.URL.Host, err)
	}
}
