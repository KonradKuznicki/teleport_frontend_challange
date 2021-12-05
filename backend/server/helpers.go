package server

import (
	"log"
	"net/http"
	"time"
)

func Handle(mux *http.ServeMux, path string, handler http.HandlerFunc) {
	log.Println("Creating path ", path)
	mux.HandleFunc(path, HSTS(LogRequests(SlowDown(handler))))
}

func HSTS(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		handler(writer, request)
	}
}

func LogRequests(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		cookie := GetCookieState(request)
		log.Printf("path: %s, cookie: %s", request.URL, cookie)
		handler(writer, request)
	}
}

func GetCookieState(request *http.Request) string {
	cookieState := "no"
	var cookieVal string
	cookie, err := request.Cookie("auth")
	if cookie != nil {
		cookieState = "yes"
		cookieVal = "val: " + cookie.Value
	}
	if err != nil {
		cookieVal += "err: " + err.Error()
	}

	return cookieState + ", " + cookieVal
}

func SlowDown(handler func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(time.Millisecond * 100)
		handler(writer, request)
	}
}
