package auth_test

import (
	"bytes"
	"encoding/json"
	"net/http"

	"challenge/auth"
)

func (s *TestAuthSuite) TestLogin_handleCorrectCreds() {
	jsonStr, _ := json.Marshal(&auth.Creds{
		Login: "user1",
		Pass:  "pass1",
	})
	req, err := http.NewRequest("POST", "/user/login", bytes.NewBuffer(jsonStr))
	if err != nil {
		s.Fail("error creating request")
	}

	http.HandlerFunc(s.Auth.LoginHandler).ServeHTTP(s.recorder, req)

	s.Equal(http.StatusFound, s.recorder.Code)
	cookie := s.recorder.Result().Cookies()[1]
	s.Equal(cookie, &http.Cookie{
		Name:     "auth",
		Value:    cookie.Value,
		Path:     "/",
		MaxAge:   2,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Raw:      cookie.Raw,
	})
}

func (s *TestAuthSuite) TestLogin_checkCredsOk() {
	_ = s.Auth.CreateUser("user1", "pass1")
	user, _ := s.Auth.GetUser("user1", "pass1")
	s.NotNil(user)
}

func (s *TestAuthSuite) TestLogin_checkCredsFail() {
	_ = s.Auth.CreateUser("user1", "pass1")
	user, _ := s.Auth.GetUser("user1", "bad pass")
	s.Nil(user)
}

func (s *TestAuthSuite) TestLogin_checkSessionFail() {
	token := "asdf"
	s.False(s.Auth.IsTokenValid(token))
}

func (s *TestAuthSuite) TestLogin_checkSessionValid() {
	_ = s.Auth.CreateUser("user1", "pass1")
	user, _ := s.Auth.GetUser("user1", "pass1")
	token, _ := s.Auth.CreateSession(user)

	s.True(s.Auth.IsTokenValid(token))
}

func (s *TestAuthSuite) TestLogin_ValidLogin() {
	_ = s.Auth.CreateUser("user1", "pass1")

	token, _ := s.Auth.Login("user1", "pass1")

	s.True(s.Auth.IsTokenValid(token))
}
