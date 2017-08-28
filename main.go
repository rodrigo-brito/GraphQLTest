package main

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/graphql-go/graphql"
	"github.com/rodrigo-brito/GraphQLTest/mutation"
	"github.com/rodrigo-brito/GraphQLTest/query"
)

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    query.UserQuery,
	Mutation: mutation.UserMutation,
})

func handle(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	result := executeQuery(query, schema)
	json.NewEncoder(w).Encode(result)
}

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main() {
	http.HandleFunc("/graphql", handle)
	fmt.Println("Now server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
