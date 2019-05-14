package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	gqlh "github.com/graphql-go/handler"
)

type GQLHandler struct {
	gqlh *gqlh.Handler
}

func NewGQLHandler(gcfg graphql.ObjectConfig) (*GQLHandler, error) {
	scfg := graphql.SchemaConfig{
		Query: graphql.NewObject(gcfg),
	}
	schema, err := graphql.NewSchema(scfg)
	if err != nil {
		return nil, err
	}

	h := &GQLHandler{
		gqlh: gqlh.New(&gqlh.Config{
			Schema:   &schema,
			Pretty:   true,
			GraphiQL: true,
		}),
	}
	return h, nil
}

func (h *GQLHandler) ContextHandler(gc *gin.Context) {
	h.gqlh.ContextHandler(gc, gc.Writer, gc.Request)
}
