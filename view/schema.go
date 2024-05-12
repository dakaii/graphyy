package view

import (
	"graphyy/controller"
	"graphyy/domain"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
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

func getSignupField(controllers *controller.Controllers) *graphql.Field {
	return &graphql.Field{
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
			res, err := controllers.UserController.Signup(domain.User{Username: username, Password: password})
			if err != nil {
				return nil, gqlerrors.FormatError(err)
			}
			return res, nil
		},
	}
}

func getLoginField(controllers *controller.Controllers) *graphql.Field {
	return &graphql.Field{
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
			res, err := controllers.UserController.Login(domain.User{Username: username, Password: password})
			if err != nil {
				return nil, gqlerrors.FormatError(err)
			}
			return res, nil
		},
	}
}

func getRootMutation(controllers *controller.Controllers) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"signup": getSignupField(controllers),
			"login":  getLoginField(controllers),
		},
	})
}

func getMeField() *graphql.Field {
	return &graphql.Field{
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
			rootValue := params.Info.RootValue.(map[string]interface{})
			user := rootValue["currentUser"].(domain.User)
			return user, nil
		},
	}
}

func getRootQuery(_ *controller.Controllers) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"me": getMeField(),
		},
	})

}
