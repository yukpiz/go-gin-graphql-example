package handler

import (
	"github.com/graphql-go/graphql"
)

var User = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id":         &graphql.Field{Type: graphql.ID},
		"first_name": &graphql.Field{Type: graphql.String},
		"last_name":  &graphql.Field{Type: graphql.String},
		"user_type":  &graphql.Field{Type: graphql.Int},
	},
})

func (uh *UserHandler) GetUserProfile() *graphql.Field {
	return &graphql.Field{
		Type: User,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			u := struct {
				ID        int    `json:"id"`
				FirstName string `json:"first_name"`
				LastName  string `json:"last_name"`
				UserType  int    `json:"user_type"`
			}{
				ID:        1,
				FirstName: "piz",
				LastName:  "yuk",
				UserType:  1,
			}
			return u, nil
		},
	}
}
