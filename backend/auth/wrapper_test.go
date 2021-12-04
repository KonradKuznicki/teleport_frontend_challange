package auth_test

import (
	"challenge/auth"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerWrapper(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(auth.Wrapper(mockHandler))

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusFound)

	// assert.Contains(t, rr.Header()., "login form")

}

func mockHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("content-type", "text/html")
	_, err := io.WriteString(writer, "<html><head></head><body><h1>file manager</h1></body></html>")
	if err != nil {
		log.Printf("error writeing response for %s error: %v", request.URL.Host, err)
	}
}
