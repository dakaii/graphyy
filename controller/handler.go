package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"graphyy/internal"
	"graphyy/repository/userrepo"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

type contextKey string

func (c contextKey) String() string {
	return "controller context key " + string(c)
}

type postData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

// BaseHandler contains all the repositories
type BaseHandler struct {
	userRepo userrepo.UserRepository
}

// NewBaseHandler returns a new BaseHandler
func NewBaseHandler(userRepo userrepo.UserRepository) *BaseHandler {
	return &BaseHandler{
		userRepo: userRepo,
	}
}

// Schema builds a graphql schema and returns it
func (h *BaseHandler) Schema() graphql.Schema {
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    h.getRootQuery(),
		Mutation: h.getRootMutation(),
	})
	return schema
}

// GraphqlHandlfunc is a handler for the graphql endpoint.
func (h *BaseHandler) GraphqlHandlfunc(w http.ResponseWriter, req *http.Request) {
	var p postData
	if err := json.NewDecoder(req.Body).Decode(&p); err != nil {
		w.WriteHeader(400)
		return
	}
	token := req.Header.Get("token")
	user, _ := internal.VerifyJWT(token)

	result := graphql.Do(graphql.Params{
		Context:        context.WithValue(context.Background(), contextKey("currentUser"), user),
		Schema:         h.Schema(),
		RequestString:  p.Query,
		VariableValues: p.Variables,
		OperationName:  p.Operation,
	})
	if len(result.Errors) > 0 {
		log.Printf("wrong result, unexpected errors: %v", result.Errors)
		return
	}
	if err := json.NewEncoder(w).Encode(result); err != nil {
		fmt.Printf("could not write result to response: %s", err)
	}
}
