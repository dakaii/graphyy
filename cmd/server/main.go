package main

import (
	"fmt"
	"graphyy/controller"
	"graphyy/database"
	"graphyy/repository"
	"graphyy/view"
	"net/http"
	"os"
)

func main() {

	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "8081"
	}

	db := database.GetDatabase()
	repos := repository.InitRepositories(db)
	controllers := controller.InitControllers(repos)
	schema := view.Schema(controllers)

	http.Handle("/graphql", view.GraphqlHandlfunc(schema))

	fmt.Println("server is started at: http://localhost:/" + port + "/")
	fmt.Println("graphql api server is started at: http://localhost:" + port + "/graphql")
	http.ListenAndServe(":"+port, nil)
}
