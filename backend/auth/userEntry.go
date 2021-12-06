package auth

import "fmt"

type UserEntry struct {
	login    string
	passHash string
}

func (u *UserEntry) PassHash() string {
	return u.passHash
}

func (u *UserEntry) ID() string {
	return u.login
}

func (u *UserEntry) String() string {
	return fmt.Sprintf(" login: %s, hash: %s", u.login, u.passHash)
}

func NewUserEntry(login string, passHash string) *UserEntry {
	return &UserEntry{
		login:    login,
		passHash: passHash,
	}
}
