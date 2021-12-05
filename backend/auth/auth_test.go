package auth_test

import (
	"challenge/auth"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestAuth(t *testing.T) {
	suite.Run(t, new(TestAuthSuite))
}

type TestAuthSuite struct {
	suite.Suite
	Auth     *auth.Auth
	handler  http.HandlerFunc
	recorder *httptest.ResponseRecorder
}

func (s *TestAuthSuite) SetupSuite() {
	s.Auth = auth.NewAuth()
}
func (s *TestAuthSuite) SetupTest() {
	s.recorder = httptest.NewRecorder()
}

func (s *TestAuthSuite) TearDownSuite() {

}

func (s *TestAuthSuite) TestAuth() {
	s.NotNil(s.Auth)
}
