package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/yukpiz/go-gin-graphql-example/handler/field"
	"github.com/yukpiz/go-gin-graphql-example/handler/resource"
)

type UserHandler struct {
	GQLHandler *GQLHandler
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (uh *UserHandler) AssignRoutes(e *gin.Engine) error {
	e.GET("/hello", func(gc *gin.Context) {
		gc.String(http.StatusOK, "Hello Redish API")
	})

	v3g := e.Group("/v3")
	{
		v3ug := v3g.Group("/user")
		{
			v3ug.GET("/hello", func(gc *gin.Context) {
				gc.String(http.StatusOK, "Hello Redish User API")
			})

			// Restful APIs
			v3ug.POST("/profile", resource.CreateUserProfile)

			// GraphQL Fields
			gcfg := graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"GetUserProfile": field.GetUserProfile(),
				},
			}

			gqlh, err := NewGQLHandler(gcfg)
			if err != nil {
				return err
			}
			v3ug.POST("/graphql", gqlh.ContextHandler)
		}
	}
	return nil
}
