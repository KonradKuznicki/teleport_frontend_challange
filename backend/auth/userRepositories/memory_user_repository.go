package userRepositories

import "challenge/auth"

type InMemoryUserRepository struct {
	usersMap map[string]auth.StorableUser
}

func (r *InMemoryUserRepository) GetUser(user auth.StorableUser) auth.StorableUser {
	return r.usersMap[user.ID()]
}

func (r *InMemoryUserRepository) AddUser(user auth.StorableUser) {
	r.usersMap[user.ID()] = user
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{usersMap: make(map[string]auth.StorableUser)}
}
