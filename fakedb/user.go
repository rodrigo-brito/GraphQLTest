package fakedb

import "github.com/rodrigo-brito/GraphQLTest/model"

type UserRepository struct{}

func (u *UserRepository) Get(ID int64) *model.User {
	for _, user := range Users {
		if user.ID == ID {
			return user
		}
	}
	return nil
}

func (u *UserRepository) All() []*model.User {
	return Users
}

func (u *UserRepository) AllValue() []model.User {
	users := []model.User{}
	for _, user := range Users {
		users = append(users, *user)
	}
	return users
}

func (u *UserRepository) Save(user *model.User) {
	Users = append(Users, user)
}

func (u *UserRepository) indexOf(ID int64) int {
	for index, user := range Users {
		if user.ID == ID {
			return index
		}
	}
	return -1
}

func (u *UserRepository) Delete(ID int64) error {
	index := u.indexOf(ID)
	if index == -1 {
		return ErrNotFound
	}
	Users = append(Users[:index], Users[index+1:]...)
	return nil
}
