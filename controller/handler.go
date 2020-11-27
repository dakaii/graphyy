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

// func (h *BaseHandler) registerAuthMutation(schema *schemabuilder.Schema) {
// 	object := schema.Mutation()
// 	object.FieldFunc("signup", func(ctx context.Context, args struct{ user *model.User }) (model.AuthToken, error) {
// 		return h.signup(*args.user)
// 	})
// 	object.FieldFunc("login", func(ctx context.Context, args struct{ user *model.User }) (model.AuthToken, error) {
// 		return h.login(*args.user)
// 	})
// }

// Schema builds a graphql schema and returns it
func (h *BaseHandler) Schema() graphql.Schema {
	// schema := schemabuilder.NewSchema()
	// h.registerAuthMutation(schema)
	// return schema.MustBuild()
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    getRootQuery(),
		Mutation: h.getRootMutation(),
	})
	return schema
}
