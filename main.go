package main

import (
	"fmt"
	"graphyy/controller"
	"graphyy/database"
	"graphyy/repository"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "8081"
	}

	db := database.GetDatabase()
	userRepo := repository.NewUserRepo(db)
	h := controller.NewBaseHandler(userRepo)

	router := mux.NewRouter()
	router.HandleFunc("/graphql", h.GraphqlHandlfunc)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	fmt.Println("Now server is running on port :" + port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
