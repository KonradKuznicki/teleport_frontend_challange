package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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
	userRepository      UserRepository
	hasher              Hasher
	sessionManager      *SessionManager
	cookieMaxAgeSeconds int
}

func (a *Auth) LoginHandler(writer http.ResponseWriter, request *http.Request) {
	creds, err := getCreds(request.Body)
	if err != nil {
		log.Println("error parsing credentials ", err.Error())
		http.Error(writer, "could not parse credentials", http.StatusBadRequest)
		return
	}
	token, err := a.Login(creds.Login, creds.Pass)
	if err != nil {
		log.Println("error logging in user ", err.Error(), creds)
		http.Error(writer, "could not authenticate", http.StatusInternalServerError)
		return
	}
	if token == "" {
		http.Error(writer, "incorrect user and or password", http.StatusForbidden)
		return
	}

	http.SetCookie(writer, &http.Cookie{
		Name:     "auth",
		Value:    token,
		MaxAge:   a.cookieMaxAgeSeconds,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
	http.Redirect(writer, request, "/files", http.StatusFound)
}

type Creds struct {
	Login string `json:"login"`
	Pass  string `json:"pass"`
}

func getCreds(body io.ReadCloser) (*Creds, error) {
	if body == nil {
		return nil, fmt.Errorf("no payload")
	}
	creds := &Creds{}
	return creds, json.NewDecoder(body).Decode(creds)
}

func (a *Auth) LogoutHandler(writer http.ResponseWriter, request *http.Request) {
	http.SetCookie(writer, &http.Cookie{
		Name:     "auth",
		Value:    "",
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

func (a *Auth) GetUser(login string, pass string) (User, error) {
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
	userEntry := NewUserEntry(login, hash)
	log.Printf("new user created: pass: %s, user: %s", pass, userEntry)
	a.userRepository.AddUser(userEntry)
	return nil
}

func (a *Auth) IsTokenValid(token string) (bool, error) {
	return a.sessionManager.Valid(token)
}

func (a *Auth) CreateSession(user StorableUser) (string, error) {
	return a.sessionManager.Create(user)
}

func (a *Auth) Login(login string, pass string) (string, error) {
	user, err := a.GetUser(login, pass)
	if err != nil {
		return "", err
	}
	return a.CreateSession(user)
}

func NewAuth(userRepository UserRepository, hasher Hasher, sessionManager *SessionManager, cookieMaxAgeSeconds int) *Auth {
	return &Auth{
		userRepository:      userRepository,
		hasher:              hasher,
		sessionManager:      sessionManager,
		cookieMaxAgeSeconds: cookieMaxAgeSeconds,
	}
}
