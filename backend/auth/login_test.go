package auth_test

import (
	"net/http"
)

func (s *TestAuthSuite) TestLogin_handleCorrectCreds() {

	req, err := http.NewRequest("GET", "/user/login", nil)
	if err != nil {
		s.Fail("error creating request")
	}

	http.HandlerFunc(s.Auth.LoginHandler).ServeHTTP(s.recorder, req)

	s.Equal(s.recorder.Code, http.StatusFound)
	cookie := s.recorder.Result().Cookies()[0]
	s.Equal(cookie, &http.Cookie{
		Name:     "auth",
		Value:    "true",
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
	user, _ := s.Auth.CheckCredsOk("user1", "pass1")
	s.NotNil(user)
}

func (s *TestAuthSuite) TestLogin_checkCredsFail() {
	_ = s.Auth.CreateUser("user1", "pass1")
	user, _ := s.Auth.CheckCredsOk("user1", "bad pass")
	s.Nil(user)
}

func (s *TestAuthSuite) TestLogin_checkSessionFail() {
	token := "asdf"
	s.False(s.Auth.CheckSessionOk(token))
}

func (s *TestAuthSuite) TestLogin_checkSessionValid() {
	_ = s.Auth.CreateUser("user1", "pass1")
	user, _ := s.Auth.CheckCredsOk("user1", "pass1")
	token, _ := s.Auth.CreateSession(user)

	s.True(s.Auth.CheckSessionOk(token))
}
