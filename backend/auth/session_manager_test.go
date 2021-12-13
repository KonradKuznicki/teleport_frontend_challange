package auth_test

import (
	"encoding/hex"
	"fmt"
	"testing"
	"time"

	"challenge/auth"
	"github.com/stretchr/testify/suite"
)

func TestSession(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(TestSessionSuite))
}

type TestSessionSuite struct {
	suite.Suite
	Session *auth.SessionManager
}

func (s *TestSessionSuite) SetupSuite() {
	s.Session = auth.NewSessionManager(&MockEncryptor{}, "some secret", time.Microsecond*100) // TODO: can be falkey better to mock time (no time though)
}

func (s *TestSessionSuite) TearDownSuite() {
}

func (s *TestSessionSuite) TestSession() {
	s.NotNil(s.Session)
}

func (s *TestSessionSuite) TestSession_NoSession() {
	s.False(s.Session.Valid("asdf"))
}

func (s *TestSessionSuite) TestSession_ValidSession() {
	token, err := s.Session.Create(auth.NewUserEntry("login", "passHash"))
	s.Nil(err)
	s.True(s.Session.Valid(token))
}

func (s *TestSessionSuite) TestSession_outdatedSession() {
	token, err := s.Session.Create(auth.NewUserEntry("login", "passHash"))
	s.Nil(err)
	time.Sleep(time.Millisecond)
	s.False(s.Session.Valid(token))
}

func (s *TestSessionSuite) TestSession_destroyedSession() {
	token, err := s.Session.Create(auth.NewUserEntry("login", "passHash"))
	s.Nil(err)
	s.Session.Destroy(token)
	s.False(s.Session.Valid(token))
}

/// mocks

type MockEncryptor struct{}

func (m MockEncryptor) Decrypt(token string) (string, error) {
	decodeString, err := hex.DecodeString(token)
	return string(decodeString), err
}

func (m MockEncryptor) Encrypt(confidential string) (string, error) {
	return fmt.Sprintf("%x", confidential), nil
}
