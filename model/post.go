package model

import "github.com/graphql-go/graphql"

var (
	postType   *graphql.Object
	nextPostID int64
)

func init() {
	initPostType()
}

type Post struct {
	ID      int64
	Title   string
	Content string
	User    *User
}

func GetPostID() int64 {
	nextPostID++
	return nextPostID
}

func initPostType() {
	postType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"content": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}
