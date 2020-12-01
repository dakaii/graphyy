package controller

import (
	"graphyy/model"

	"github.com/graphql-go/graphql"
)

var authType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Auth",
	Fields: graphql.Fields{
		"tokenType": &graphql.Field{
			Type: graphql.String,
		},
		"token": &graphql.Field{
			Type: graphql.String,
		},
		"expiresIn": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

func (h *BaseHandler) getRootMutation() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"signup": &graphql.Field{
				Type:        authType, // the return type for this field
				Description: "Signup",
				Args: graphql.FieldConfigArgument{
					"username": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					username, _ := params.Args["username"].(string)
					password, _ := params.Args["password"].(string)
					return h.signup(model.User{Username: username, Password: password})
				},
			},
			"login": &graphql.Field{
				Type:        authType, // the return type for this field
				Description: "Login",
				Args: graphql.FieldConfigArgument{
					"username": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					username, _ := params.Args["username"].(string)
					password, _ := params.Args["password"].(string)
					return h.login(model.User{Username: username, Password: password})
				},
			},
		},
	})
}

func (h *BaseHandler) getRootQuery() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"me": &graphql.Field{
				Type: graphql.NewObject(graphql.ObjectConfig{
					Name: "Me",
					Fields: graphql.Fields{
						"username": &graphql.Field{
							Type: graphql.String,
						},
					},
				}),
				Description: "Get the logged-in user's info",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					user := params.Context.Value("currentUser").(model.User)
					return user.Username, nil
				},
			},
		},
	})

}
