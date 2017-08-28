package model

import "github.com/graphql-go/graphql"

var (
	nextUserID int64 = 0
	UserType   *graphql.Object
)

func init() {
	initUserType()
}

func GetUserID() int64 {
	nextUserID++
	return nextUserID
}

type User struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Posts []*Post `json:"posts"`
}

func initUserType() {
	UserType = graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}
