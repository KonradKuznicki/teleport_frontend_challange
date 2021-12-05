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
		MaxAge:   300,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Raw:      cookie.Raw,
	})
}
