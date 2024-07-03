package main

import (
	"fmt"
	"net/http"

	"github.com/dakaii/graphyy/internal/controller"
	"github.com/dakaii/graphyy/internal/database"
	"github.com/dakaii/graphyy/internal/envvar"
	"github.com/dakaii/graphyy/internal/repository"
	"github.com/dakaii/graphyy/internal/view"
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
