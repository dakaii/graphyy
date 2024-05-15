package main

import (
	"fmt"
	"graphyy/controller"
	"graphyy/database"
	"graphyy/internal/envvar"
	"graphyy/repository"
	"graphyy/view"
	"net/http"
)

func main() {

	db := database.GetDatabase()
	repos := repository.InitRepositories(db)
	controllers := controller.InitControllers(repos)
	schema := view.Schema(controllers)

	http.Handle("/graphql", view.GraphqlHandlfunc(schema))

	port := envvar.Port()
	fmt.Println("server is started at: http://localhost:/" + port + "/")
	fmt.Println("graphql api server is started at: http://localhost:" + port + "/graphql")
	http.ListenAndServe(":"+port, nil)
}
