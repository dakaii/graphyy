package main

import (
	"graphyy/controller"
	"graphyy/database"
	"graphyy/repository"
	"net/http"
	"os"

	"github.com/samsarahq/thunder/graphql"

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

	http.Handle("/graphql", graphql.HTTPHandler(schema))
	http.ListenAndServe(":"+port, nil)
}
