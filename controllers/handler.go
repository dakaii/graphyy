package controllers

import (
	"coldhongdae/models"

	"github.com/samsarahq/thunder/graphql"
	"github.com/samsarahq/thunder/graphql/schemabuilder"
)

// BaseHandler contains all the repositories
type BaseHandler struct {
	userRepo models.UserRepository
}

// NewBaseHandler returns a new BaseHandler
func NewBaseHandler(userRepo models.UserRepository) *BaseHandler {
	return &BaseHandler{
		userRepo: userRepo,
	}
}

func (h *BaseHandler) registerAuthMutation(schema *schemabuilder.Schema) {
	object := schema.Mutation()
	object.FieldFunc("signup", h.Signup)
	object.FieldFunc("login", h.Login)
}

// Schema builds a graphql schema and returns it
func (h *BaseHandler) Schema() *graphql.Schema {
	builder := schemabuilder.NewSchema()
	h.registerAuthMutation(builder)
	return builder.MustBuild()
}
