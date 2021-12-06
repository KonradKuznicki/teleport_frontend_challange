package auth_test

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/argon2"
)

type HardHasher struct {
	format  string
	version int
	time    uint32
	memory  uint32
	keyLen  uint32
	saltLen uint32
	threads uint8
}

func (h *HardHasher) Hash(pass string) (string, error) {
	salt := make([]byte, h.saltLen)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(pass), salt, h.time, h.memory, h.threads, h.keyLen)

	return fmt.Sprintf(
			h.format,
			h.version,
			h.memory,
			h.time,
			h.threads,
			base64.RawStdEncoding.EncodeToString(salt),
			base64.RawStdEncoding.EncodeToString(hash),
		),
		nil
}

func (h *HardHasher) Verify(plain, hash string) (bool, error) {
	hashParts := strings.Split(hash, "$")

	_, err := fmt.Sscanf(hashParts[3], "m=%d,t=%d,p=%d", &h.memory, &h.time, &h.threads)
	if err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(hashParts[4])
	if err != nil {
		return false, err
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(hashParts[5])
	if err != nil {
		return false, err
	}

	hashToCompare := argon2.IDKey([]byte(plain), salt, h.time, h.memory, h.threads, uint32(len(decodedHash)))

	return subtle.ConstantTimeCompare(decodedHash, hashToCompare) == 1, nil
}

func NewHardHasher() *HardHasher {
	return &HardHasher{
		format:  "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		version: argon2.Version,
		time:    1,
		memory:  64 * 1024,
		keyLen:  32,
		saltLen: 16,
		threads: 4}
}

func TestHardHasher(t *testing.T) {
	suite.Run(t, new(TestHardHasherSuite))
}

type TestHardHasherSuite struct {
	suite.Suite
	HardHasher *HardHasher
}

func (s *TestHardHasherSuite) SetupSuite() {
	s.HardHasher = NewHardHasher()
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
