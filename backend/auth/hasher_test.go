package auth

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestEasyHash(t *testing.T) {
	suite.Run(t, new(TestEasyHashSuite))
}

type TestEasyHashSuite struct {
	suite.Suite
	EasyHash *EasyHash
}

func (s *TestEasyHashSuite) SetupSuite() {
	s.EasyHash = NewEasyHash("salt")
}

func (s *TestEasyHashSuite) TearDownSuite() {

}

func (s *TestEasyHashSuite) TestEasyHash() {
	s.NotNil(s.EasyHash)
}

func (s *TestEasyHashSuite) TestEasyHash_Hash() {
	hash, _ := s.EasyHash.Hash("asdf")
	s.Equal(hash, "asdfsalt")
}
