package controller

import (
	"context"
	"graphyy/internal"
	"log"
	"net/http"
	"strings"

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
			authorization := req.Header.Get("Authorization")
			if authorization == "" {
				return map[string]interface{}{
					"currentUser": nil,
				}
			}

			const bearerPrefix = "Bearer "
			if !strings.HasPrefix(authorization, bearerPrefix) {
				log.Println("Invalid authorization format")
				return map[string]interface{}{
					"currentUser": nil,
				}
			}

			token := strings.TrimPrefix(authorization, bearerPrefix)
			user, err := internal.VerifyJWT(token)
			if err != nil {
				log.Printf("Failed to verify JWT: %v", err)
				return map[string]interface{}{
					"currentUser": nil,
				}
			}

			return map[string]interface{}{
				"currentUser": user,
			}
		},
	})
}
