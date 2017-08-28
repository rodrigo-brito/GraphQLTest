package fakedb

import (
	"errors"

	"github.com/rodrigo-brito/GraphQLTest/model"
)

var (
	ErrNotFound = errors.New("Not found.")
	Posts       []*model.Post
	Users       []*model.User
)

func init() {
	Users = []*model.User{
		&model.User{ID: model.GetUserID(), Name: "João"},
		&model.User{ID: model.GetUserID(), Name: "Maria"},
		&model.User{ID: model.GetUserID(), Name: "Pedro"},
		&model.User{ID: model.GetUserID(), Name: "Joaquim"},
	}
	Posts = []*model.Post{
		&model.Post{ID: model.GetPostID(), Title: "Post 1", Content: "Content 1", User: Users[0]},
		&model.Post{ID: model.GetPostID(), Title: "Post 2", Content: "Content 2", User: Users[3]},
		&model.Post{ID: model.GetPostID(), Title: "Post 3", Content: "Content 3", User: Users[1]},
		&model.Post{ID: model.GetPostID(), Title: "Post 4", Content: "Content 4", User: Users[0]},
		&model.Post{ID: model.GetPostID(), Title: "Post 5", Content: "Content 5", User: Users[2]},
		&model.Post{ID: model.GetPostID(), Title: "Post 6", Content: "Content 6", User: Users[1]},
		&model.Post{ID: model.GetPostID(), Title: "Post 7", Content: "Content 7", User: Users[2]},
		&model.Post{ID: model.GetPostID(), Title: "Post 8", Content: "Content 8", User: Users[3]},
	}
}
