package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yukpiz/go-gin-graphql-example/handler"
)

func main() {
	e, err := NewGin()
	if err != nil {
		panic(err)
	}

	if err := AssignRoutes(e); err != nil {
		panic(err)
	}

	http.Handle("/", e)
	http.ListenAndServe(":8888", nil)
}

func NewGin() (*gin.Engine, error) {
	e := gin.Default()
	return e, nil
}

func AssignRoutes(e *gin.Engine) error {
	uh := handler.NewUserHandler()
	if err := uh.AssignRoutes(e); err != nil {
		return err
	}

	return nil
}
