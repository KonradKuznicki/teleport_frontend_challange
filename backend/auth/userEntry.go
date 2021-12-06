package auth

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

func NewUserEntry(login string, passHash string) *UserEntry {
	return &UserEntry{
		login:    login,
		passHash: passHash,
	}
}
