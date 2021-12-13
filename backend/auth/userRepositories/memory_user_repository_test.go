package userRepositories_test

import (
	"testing"

	"challenge/auth"
	"challenge/auth/userRepositories"
	"github.com/stretchr/testify/suite"
)

func TestInMemoryUserRepository(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(TestInMemoryUserRepositorySuite))
}

type TestInMemoryUserRepositorySuite struct {
	suite.Suite
	InMemoryUserRepository *userRepositories.InMemoryUserRepository
}

func (s *TestInMemoryUserRepositorySuite) SetupSuite() {
	s.InMemoryUserRepository = userRepositories.NewInMemoryUserRepository()
}

func (s *TestInMemoryUserRepositorySuite) TearDownSuite() {
}

func (s *TestInMemoryUserRepositorySuite) TestInMemoryUserRepository() {
	s.NotNil(s.InMemoryUserRepository)
}

func (s *TestInMemoryUserRepositorySuite) TestInMemoryUserRepository_userDoesntExist() {
	user := s.InMemoryUserRepository.GetUser(auth.NewUserQuery("pad id"))
	s.Nil(user)
}

func (s *TestInMemoryUserRepositorySuite) TestInMemoryUserRepository_userExists() {
	s.InMemoryUserRepository.AddUser(auth.NewUserEntry("user1", "asdf"))
	user := s.InMemoryUserRepository.GetUser(auth.NewUserQuery("user1"))
	s.NotNil(user)
}
