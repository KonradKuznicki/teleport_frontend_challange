package auth

import (
	"fmt"
	"net/http"
	"time"
)

type StorableUser interface {
	ID() string
}

type User interface {
	ID() string
	PassHash() string
}

type UserCreator func(login string, passHash string) User

type UserRepository interface {
	GetUser(user StorableUser) StorableUser
	AddUser(user StorableUser)
}

type Hasher interface {
	Hash(string string) (string, error)
	Verify(pass string, hash string) (bool, error)
}

type Auth struct {
	userRepository UserRepository
	hasher         Hasher
}

func (a *Auth) LoginHandler(writer http.ResponseWriter, request *http.Request) {
	http.SetCookie(writer, &http.Cookie{
		Name:     "auth",
		Value:    "true",
		MaxAge:   300,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
	http.Redirect(writer, request, "/files", http.StatusFound)
}

func (a *Auth) LogoutHandler(writer http.ResponseWriter, request *http.Request) {
	http.SetCookie(writer, &http.Cookie{
		Name:     "auth",
		Value:    "false",
		Expires:  time.Unix(0, 0),
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
	http.Redirect(writer, request, "/login", http.StatusFound)
}

func NewUserQuery(id string) *UserQuery {
	return &UserQuery{id: id}
}

type UserQuery struct {
	id string
}

func (u *UserQuery) ID() string {
	return u.id
}

func (a *Auth) CheckCredsOk(login string, pass string) (User, error) {
	storableUser := a.userRepository.GetUser(&UserQuery{id: login})
	if storableUser == nil {
		return nil, nil
	}

	user, ok := storableUser.(User)
	if !ok {
		return nil, fmt.Errorf("can't convert stored user back to user")
	}

	verify, err := a.hasher.Verify(pass, user.PassHash())
	if err != nil {
		return nil, err
	}

	if verify {
		return user, nil
	}

	return nil, nil
}

func (a *Auth) CreateUser(login string, pass string) error {
	hash, err := a.hasher.Hash(pass)

	if err != nil {
		return err
	}
	a.userRepository.AddUser(NewUserEntry(login, hash))
	return nil
}

func NewAuth(userRepository UserRepository, hasher Hasher) *Auth {
	return &Auth{
		userRepository: userRepository,
		hasher:         hasher,
	}
}
