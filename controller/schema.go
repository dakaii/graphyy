package controller

import (
	"graphyy/model"

	"github.com/graphql-go/graphql"
)

// var TodoList []Todo

// type Todo struct {
// 	ID   string `json:"id"`
// 	Text string `json:"text"`
// 	Done bool   `json:"done"`
// }

// var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// func RandStringRunes(n int) string {
// 	b := make([]rune, n)
// 	for i := range b {
// 		b[i] = letterRunes[rand.Intn(len(letterRunes))]
// 	}
// 	return string(b)
// }

// define custom GraphQL ObjectType `todoType` for our Golang struct `Todo`
// Note that
// - the fields in our todoType maps with the json tags for the fields in our struct
// - the field type matches the field type in our struct
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

// root query
// we just define a trivial example here, since root query is required.
// Test with curl
// curl -g 'http://localhost:8080/graphql?query={lastTodo{id,text,done}}'
// func getRootQuery() *graphql.Object {
// 	return graphql.NewObject(graphql.ObjectConfig{
// 		Name: "RootQuery",
// 		Fields: graphql.Fields{

// 			/*
// 			   curl -g 'http://localhost:8080/graphql?query={todo(id:"b"){id,text,done}}'
// 			*/
// 			"todo": &graphql.Field{
// 				Type:        todoType,
// 				Description: "Get single todo",
// 				Args: graphql.FieldConfigArgument{
// 					"id": &graphql.ArgumentConfig{
// 						Type: graphql.String,
// 					},
// 				},
// 				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

// 					idQuery, isOK := params.Args["id"].(string)
// 					if isOK {
// 						// Search for el with id
// 						for _, todo := range TodoList {
// 							if todo.ID == idQuery {
// 								return todo, nil
// 							}
// 						}
// 					}

// 					return Todo{}, nil
// 				},
// 			},

// 			"lastTodo": &graphql.Field{
// 				Type:        todoType,
// 				Description: "Last todo added",
// 				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
// 					return TodoList[len(TodoList)-1], nil
// 				},
// 			},

// 			/*
// 			   curl -g 'http://localhost:8080/graphql?query={todoList{id,text,done}}'
// 			*/
// 			"todoList": &graphql.Field{
// 				Type:        graphql.NewList(todoType),
// 				Description: "List of todos",
// 				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 					return TodoList, nil
// 				},
// 			},
// 		},
// 	})

// }

func getRootQuery() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"todo": &graphql.Field{
				Type: graphql.NewObject(graphql.ObjectConfig{
					Name: "Todo",
					Fields: graphql.Fields{
						"todo": &graphql.Field{
							Type: graphql.String,
						},
					},
				}),
				Description: "Get single todo",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(string)
					return id, nil
				},
			},
		},
	})

}
