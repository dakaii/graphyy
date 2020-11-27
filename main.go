package main

import (
	"encoding/json"
	"fmt"
	"graphyy/controller"
	"graphyy/database"
	"graphyy/repository"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
)

// type postData struct {
// 	Query     string                 `json:"query"`
// 	Operation string                 `json:"operation"`
// 	Variables map[string]interface{} `json:"variables"`
// }

func main() {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "8081"
	}

	db := database.GetDatabase()
	userRepo := repository.NewUserRepo(db)
	h := controller.NewBaseHandler(userRepo)

	// schema := h.Schema()
	// introspection.AddIntrospectionToSchema(schema)

	// http.Handle("/graphql", graphql.HTTPHandler(schema))
	// http.ListenAndServe(":"+port, nil)
	// http.HandleFunc("/graphql", func(w http.ResponseWriter, req *http.Request) {
	// 	var p postData
	// 	if err := json.NewDecoder(req.Body).Decode(&p); err != nil {
	// 		w.WriteHeader(400)
	// 		return
	// 	}
	// 	result := graphql.Do(graphql.Params{
	// 		Context:        req.Context(),
	// 		Schema:         h.Schema(),
	// 		RequestString:  p.Query,
	// 		VariableValues: p.Variables,
	// 		OperationName:  p.Operation,
	// 	})
	// 	if err := json.NewEncoder(w).Encode(result); err != nil {
	// 		fmt.Printf("could not write result to response: %s", err)
	// 	}
	// })
	router := mux.NewRouter()
	//api route is /people,
	//Methods("GET", "OPTIONS") means it support GET, OPTIONS
	router.HandleFunc("/graphql", func(w http.ResponseWriter, req *http.Request) {
		query := req.URL.Query().Get("query")
		result := graphql.Do(graphql.Params{
			Schema:        h.Schema(),
			RequestString: query,
		})
		json.NewEncoder(w).Encode(result)
		// var p postData
		// if err := json.NewDecoder(req.Body).Decode(&p); err != nil {
		// 	w.WriteHeader(400)
		// 	return
		// }
		// result := graphql.Do(graphql.Params{
		// 	Context:        req.Context(),
		// 	Schema:         h.Schema(),
		// 	RequestString:  p.Query,
		// 	VariableValues: p.Variables,
		// 	OperationName:  p.Operation,
		// })
		// if err := json.NewEncoder(w).Encode(result); err != nil {
		// 	fmt.Printf("could not write result to response: %s", err)
		// }
	})
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	fmt.Println("Now server is running on port :" + port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
