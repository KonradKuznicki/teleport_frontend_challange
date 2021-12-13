package auth_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"challenge/auth"
	"challenge/auth/userRepositories"
	"github.com/stretchr/testify/suite"
)

func TestAuth(t *testing.T) {
	t.Parallel()
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
	s.hasher = auth.NewEasyHash("salt")
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
