package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserProfile(gc *gin.Context) {
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
	gc.JSON(http.StatusOK, u)
}
