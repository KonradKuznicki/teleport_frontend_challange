package auth_test

import (
	"testing"

	"challenge/auth"
	"github.com/stretchr/testify/suite"
)

func TestHardHasher(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(TestHardHasherSuite))
}

type TestHardHasherSuite struct {
	suite.Suite
	HardHasher *auth.HardHasher
}

func (s *TestHardHasherSuite) SetupSuite() {
	s.HardHasher = auth.NewHardHasher()
}

func (s *TestHardHasherSuite) TearDownSuite() {
}

func (s *TestHardHasherSuite) TestHardHasher() {
	s.NotNil(s.HardHasher)
}

func (s *TestHardHasherSuite) TestEasyHash_VerifyTrue() {
	pass := "pass"
	hash, err := s.HardHasher.Hash(pass)
	s.Nil(err)
	s.True(s.HardHasher.Verify(pass, hash))
}

func (s *TestHardHasherSuite) TestEasyHash_VerifyFalse() {
	hash, err := s.HardHasher.Hash("good pass")
	s.Nil(err)
	s.False(s.HardHasher.Verify("bad pass", hash))
}
