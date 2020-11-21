package main

import (
	"graphyy/controller"
	"graphyy/database"
	"graphyy/repository"
	"net/http"
	"os"

	"github.com/samsarahq/thunder/graphql"
	"github.com/samsarahq/thunder/graphql/graphiql"
	"github.com/samsarahq/thunder/graphql/introspection"
)

func main() {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "3030"
	}

	db := database.GetDatabase()
	userRepo := repository.NewUserRepo(db)
	h := controller.NewBaseHandler(userRepo)

	schema := h.Schema()
	introspection.AddIntrospectionToSchema(schema)

	// Expose schema and graphiql.
	http.Handle("/graphql", graphql.Handler(schema))
	http.Handle("/graphiql/", http.StripPrefix("/graphiql/", graphiql.Handler()))
	http.ListenAndServe(":"+port, nil)
}
