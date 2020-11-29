package controller

import (
	"graphyy/repository"

	// "github.com/samsarahq/thunder/graphql"
	"github.com/graphql-go/graphql"
)

// BaseHandler contains all the repositories
type BaseHandler struct {
	userRepo repository.UserRepository
}

// NewBaseHandler returns a new BaseHandler
func NewBaseHandler(userRepo repository.UserRepository) *BaseHandler {
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
