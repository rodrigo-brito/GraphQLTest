package main

import (
	"net/http"

	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/rodrigo-brito/GraphQLTest/model"
)

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    model.UserQuery,
	Mutation: model.UserMutation,
})

func main() {
	handle := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	fs := http.FileServer(http.Dir("static"))

	http.Handle("/graphql", handle)
	http.Handle("/", fs)
	fmt.Printf("Server started at: localhost:8080")
	http.ListenAndServe(":8080", nil)
}
