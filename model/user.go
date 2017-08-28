package model

import (
	"github.com/graphql-go/graphql"
	"github.com/rodrigo-brito/GraphQLTest/fakedb"
)

var (
	UserType     *graphql.Object
	UserMutation *graphql.Object
	UserQuery    *graphql.Object
)

func init() {
	initUserType()
}

func initUserType() {
	UserType = graphql.NewObject(graphql.ObjectConfig{
		Name: "UserType",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"posts": &graphql.Field{
				Type:        graphql.NewList(PostType),
				Description: "List of user's post",
				Args: graphql.FieldConfigArgument{
					"limit": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					if user, ok := params.Source.(*fakedb.User); ok {
						repo := fakedb.PostRepository{}
						posts := repo.ByUserID(user.ID)
						limit := params.Args["limit"].(int)
						if len(posts) > limit {
							return posts[:limit], nil
						}
						return posts, nil
					}
					return nil, nil
				},
			},
		},
	})

	UserQuery = graphql.NewObject(graphql.ObjectConfig{
		Name: "UserQuery",
		Fields: graphql.Fields{
			"users": &graphql.Field{
				Type:        graphql.NewList(UserType),
				Description: "List of users",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					repo := fakedb.UserRepository{}
					return repo.All(), nil
				},
			},
		},
	})

	UserMutation = graphql.NewObject(graphql.ObjectConfig{
		Name: "UserMutation",
		Fields: graphql.Fields{
			"createUser": &graphql.Field{
				Type:        UserType,
				Description: "Create new user",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					name, _ := params.Args["name"].(string)

					newUser := fakedb.User{
						ID:   fakedb.GenerateUserID(),
						Name: name,
					}
					repo := fakedb.UserRepository{}
					repo.Save(&newUser)

					return newUser, nil
				},
			},
		},
	})
}
