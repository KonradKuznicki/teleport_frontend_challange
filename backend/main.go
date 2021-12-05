package main

import (
	"challenge/auth"
	"challenge/files"
	"crypto/tls"
	"log"
	"net/http"
	"time"
)

func main() {
	// Set routing rules

	log.Println("starting")

	mux := http.NewServeMux()
	//mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	//	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	//	w.Write([]byte("This is an example server.\n"))
	//})
	authenticator := auth.NewAuth()
	// http.HandleFunc("/user/logout", slowDown(authenticator.LogoutHandler))
	handle(mux, "/user/logout", authenticator.LogoutHandler)
	handle(mux, "/user/login", authenticator.LoginHandler)
	handle(mux, "/login", auth.StaticsHandler)
	handle(mux, "/files", auth.Wrapper(files.SataticsHandler))
	handle(mux, "/", IndexHandler)

	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{

			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, // Go 1.8 only
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,   // Go 1.8 only
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,

			//tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			//tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
		},
	}
	srv := &http.Server{
		Addr:         ":3001",
		Handler:      mux,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	err := srv.ListenAndServeTLS("server.ecdsa.crt", "server.ecdsa.key")
	if err != nil {
		log.Println("errrr: " + err.Error())
	}

	//Use the default DefaultServeMux.
	//err = http.ListenAndServe(":3000", nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
}

func handle(mux *http.ServeMux, path string, handler func(writer http.ResponseWriter, request *http.Request)) {

	mux.HandleFunc(path, hsts(logHttp(slowDown(handler))))
}

func hsts(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		f(writer, request)
	}
}

func logHttp(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		cookie := getCookieState(request)
		log.Printf("path: %s, cookie: %s", request.URL, cookie)
		handler(writer, request)
	}
}

func getCookieState(request *http.Request) string {
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

func slowDown(handler func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(time.Millisecond * 100)
		handler(writer, request)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/files", http.StatusPermanentRedirect)
	} else {
		http.NotFound(w, r)
	}
}
