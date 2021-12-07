package files

import (
	"net/http"
)

func SataticsHandler(writer http.ResponseWriter, request *http.Request) {
	fs := http.FileServer(http.Dir("../frontend/build"))
	http.StripPrefix("/files", fs).ServeHTTP(writer, request)
}

//writer.Header().Add("content-type", "text/html")
//_, err := io.WriteString(writer, "<html><head></head><body><h1>files manager</h1></body></html>")
//if err != nil {
//log.Printf("error writeing response for %s error: %v", request.URL.Host, err)
//}
