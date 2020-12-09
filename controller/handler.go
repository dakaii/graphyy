package controller

import (
	"context"
	"graphyy/internal"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// Schema builds a graphql schema and returns it
func Schema(controllers *Controllers) graphql.Schema {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    getRootQuery(controllers),
		Mutation: getRootMutation(controllers),
	})
	if err != nil {
		panic(err)
	}

	return schema
}

// GraphqlHandlfunc is a handler for the graphql endpoint.
func GraphqlHandlfunc(schema graphql.Schema) *handler.Handler {
	return handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
		RootObjectFn: func(ctx context.Context, req *http.Request) map[string]interface{} {
			token := req.Header.Get("token")
			user, _ := internal.VerifyJWT(token)
			return map[string]interface{}{
				"currentUser": user,
			}
		},
	})
}
