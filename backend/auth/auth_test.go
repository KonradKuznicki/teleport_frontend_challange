package auth_test

import (
	"challenge/auth"
	"challenge/auth/userRepositories"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

func TestAuth(t *testing.T) {
	suite.Run(t, new(TestAuthSuite))
}

type TestAuthSuite struct {
	suite.Suite
	Auth                   *auth.Auth
	handler                http.HandlerFunc
	recorder               *httptest.ResponseRecorder
	InMemoryUserRepository *userRepositories.InMemoryUserRepository
	hasher                 *auth.EasyHash
}

func (s *TestAuthSuite) SetupSuite() {
	s.InMemoryUserRepository = userRepositories.NewInMemoryUserRepository()
	s.hasher = auth.NewEasyHash("aslt")
	s.Auth = auth.NewAuth(s.InMemoryUserRepository, s.hasher, auth.NewSessionManager(&MockEncryptor{}, "magic secret", time.Second), 2)
}
func (s *TestAuthSuite) SetupTest() {
	s.recorder = httptest.NewRecorder()
}

func (s *TestAuthSuite) TearDownSuite() {

}

func (s *TestAuthSuite) TestAuth() {
	s.NotNil(s.Auth)
}
