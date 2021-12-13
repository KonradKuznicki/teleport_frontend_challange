package auth_test

import (
	"fmt"
	"strings"
	"testing"

	"challenge/auth"
	"github.com/stretchr/testify/suite"
)

func TestAESEncryptor(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(TestAESEncryptorSuite))
}

type TestAESEncryptorSuite struct {
	suite.Suite
	AESEncryptor *auth.AESEncryptor
}

func (s *TestAESEncryptorSuite) SetupSuite() {
	s.AESEncryptor = auth.NewAESEncryptor(fmt.Sprintf("%x", strings.Repeat("A", 32)))
}

func (s *TestAESEncryptorSuite) TearDownSuite() {
}

func (s *TestAESEncryptorSuite) TestAESEncryptor() {
	s.NotNil(s.AESEncryptor)
}

func (s *TestAESEncryptorSuite) TestAESEncryptor_ok() {
	token := "asdf"
	encrypt, err := s.AESEncryptor.Encrypt(token)
	s.Nil(err)
	decrypt, err := s.AESEncryptor.Decrypt(encrypt)
	s.Nil(err)
	s.Equal(decrypt, token)
}
