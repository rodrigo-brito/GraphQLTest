package model

import "github.com/graphql-go/graphql"

var (
	PostType *graphql.Object
)

func init() {
	initPostType()
}

func initPostType() {
	PostType = graphql.NewObject(graphql.ObjectConfig{
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
