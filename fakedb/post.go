package fakedb

import (
	"github.com/rodrigo-brito/GraphQLTest/model"
)

type PostRepository struct{}

func (p *PostRepository) Get(ID int64) *model.Post {
	for _, post := range Posts {
		if post.ID == ID {
			return post
		}
	}
	return nil
}

func (p *PostRepository) All() []*model.Post {
	return Posts
}

func (p *PostRepository) Save(post *model.Post) {
	Posts = append(Posts, post)
}

func (p *PostRepository) indexOf(ID int64) int {
	for index, post := range Posts {
		if post.ID == ID {
			return index
		}
	}
	return -1
}

func (p *PostRepository) Delete(ID int64) error {
	index := p.indexOf(ID)
	if index == -1 {
		return ErrNotFound
	}
	Posts = append(Posts[:index], Posts[index+1:]...)
	return nil
}
