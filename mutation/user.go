package mutation

import (
	"github.com/graphql-go/graphql"
	"github.com/rodrigo-brito/GraphQLTest/fakedb"
	"github.com/rodrigo-brito/GraphQLTest/model"
)

var UserMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "AddUser",
	Fields: graphql.Fields{
		"createUser": &graphql.Field{
			Type:        model.UserType, // the return type for this field
			Description: "Create new user",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				name, _ := params.Args["name"].(string)

				newUser := model.User{
					ID:   model.GetUserID(),
					Name: name,
				}
				repo := fakedb.UserRepository{}
				repo.Save(&newUser)

				return newUser, nil
			},
		},
	},
})
