package fakedb

import (
	"errors"
)

var (
	ErrNotFound = errors.New("Not found.")
	Posts       []*Post
	Users       []*User
)

func init() {
	Users = []*User{
		&User{ID: GenerateUserID(), Name: "João"},
		&User{ID: GenerateUserID(), Name: "Maria"},
		&User{ID: GenerateUserID(), Name: "Pedro"},
		&User{ID: GenerateUserID(), Name: "Joaquim"},
	}
	Posts = []*Post{
		&Post{ID: GeneratePostID(), Title: "Post 1", Content: "Content 1", User: Users[0]},
		&Post{ID: GeneratePostID(), Title: "Post 2", Content: "Content 2", User: Users[3]},
		&Post{ID: GeneratePostID(), Title: "Post 3", Content: "Content 3", User: Users[1]},
		&Post{ID: GeneratePostID(), Title: "Post 4", Content: "Content 4", User: Users[0]},
		&Post{ID: GeneratePostID(), Title: "Post 5", Content: "Content 5", User: Users[2]},
		&Post{ID: GeneratePostID(), Title: "Post 6", Content: "Content 6", User: Users[1]},
		&Post{ID: GeneratePostID(), Title: "Post 7", Content: "Content 7", User: Users[2]},
		&Post{ID: GeneratePostID(), Title: "Post 8", Content: "Content 8", User: Users[3]},
		&Post{ID: GeneratePostID(), Title: "Post 9", Content: "Content 9", User: Users[3]},
		&Post{ID: GeneratePostID(), Title: "Post 10", Content: "Content 10", User: Users[3]},
		&Post{ID: GeneratePostID(), Title: "Post 11", Content: "Content 11", User: Users[3]},
		&Post{ID: GeneratePostID(), Title: "Post 12", Content: "Content 12", User: Users[3]},
		&Post{ID: GeneratePostID(), Title: "Post 13", Content: "Content 13", User: Users[3]},
	}
}
