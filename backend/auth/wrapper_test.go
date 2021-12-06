package auth_test

import (
	"io"
	"log"
	"net/http"
)

func (s *TestAuthSuite) TestWrapper_redirectWhenNotAuthenticated() {
	s.handler = s.Auth.Wrapper(mockHandler)

	s.NotNil(s.Auth)
	req, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		s.Fail("error creating request")
	}

	s.handler.ServeHTTP(s.recorder, req)

	s.Equal(s.recorder.Code, http.StatusFound)
}

func (s *TestAuthSuite) TestWrapper_dontRedirectWhenAuthenticated() {
	_ = s.Auth.CreateUser("user1", "pass1")
	token, _ := s.Auth.Login("user1", "pass1")
	s.handler = s.Auth.Wrapper(mockHandler)

	s.NotNil(s.Auth)
	req, err := http.NewRequest("GET", "/files", nil)
	if err != nil {
		s.Fail("error creating request")
	}

	http.SetCookie(s.recorder, &http.Cookie{Name: "auth", Value: token})
	req.Header.Set("Cookie", s.recorder.Header().Get("Set-Cookie"))

	s.handler.ServeHTTP(s.recorder, req)

	s.Equal(s.recorder.Code, http.StatusOK)
}

////// mocks

func mockHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("content-type", "text/html")
	_, err := io.WriteString(writer, "<html><head></head><body><h1>file manager</h1></body></html>")
	if err != nil {
		log.Printf("error writeing response for %s error: %v", request.URL.Host, err)
	}
}
