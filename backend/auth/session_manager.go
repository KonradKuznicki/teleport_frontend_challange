package auth

import (
	"encoding/json"
	"log"
	"sync"
	"time"
)

type Encryptor interface {
	Decrypt(token string) (string, error)
	Encrypt(confidential string) (string, error)
}

type SessionManager struct {
	mu            sync.Mutex
	invalidTokens map[string]time.Time
	secret        string
	encryptor     Encryptor
	maxAge        time.Duration
}

func (m *SessionManager) Valid(opaqueToken string) (bool, error) {
	if m.isInvalidated(opaqueToken) {
		return false, nil
	}
	decrypted, _ := m.encryptor.Decrypt(opaqueToken)

	sess, err := UnmarshalSession(decrypted)
	if err != nil {
		log.Println(err)
		return false, err
	}
	if sess.Secret != m.secret {
		log.Println("bad secret")
		return false, nil
	}
	if !sess.Time.After(time.Now().Add(-m.maxAge)) {
		log.Println("bad time", sess.Time, time.Now().Add(-m.maxAge))
		return false, nil
	}
	return true, nil
}

func (m *SessionManager) Create(user StorableUser) (string, error) {
	session := NewSession(m.secret, time.Now(), user.ID())
	data, err := json.Marshal(session)
	if err != nil {
		return "", err
	}
	return m.encryptor.Encrypt(string(data))
}

func (m *SessionManager) Destroy(token string) {
	m.invalidate(token)
}

func (m *SessionManager) invalidate(token string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.invalidTokens[token] = time.Now() // TODO: garbage collect
}

func (m *SessionManager) isInvalidated(token string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	_, ok := m.invalidTokens[token]
	return ok
}

func NewSessionManager(encryptor Encryptor, secret string, maxAge time.Duration) *SessionManager {
	return &SessionManager{invalidTokens: make(map[string]time.Time), encryptor: encryptor, secret: secret, maxAge: maxAge}
}

type session struct {
	Secret string
	Time   time.Time
	UserID string
}

func NewSession(secret string, time time.Time, userID string) *session {
	return &session{Secret: secret, Time: time, UserID: userID}
}

func UnmarshalSession(data string) (*session, error) {
	sess := &session{}
	err := json.Unmarshal([]byte(data), sess)
	if err != nil {
		return nil, err
	}
	return sess, nil
}
