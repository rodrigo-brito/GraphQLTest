package query

import (
	"github.com/graphql-go/graphql"
	"github.com/rodrigo-brito/GraphQLTest/fakedb"
	"github.com/rodrigo-brito/GraphQLTest/model"
)

var UserQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"user": &graphql.Field{
			Type:        graphql.NewList(model.UserType),
			Description: "List of users",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				repo := fakedb.UserRepository{}
				return repo.All(), nil
			},
		},
	},
})
