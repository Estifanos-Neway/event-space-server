package handlers

import (
	"net/http"

	"github.com/estifanos-neway/event-space-server/src/repos"
	"github.com/estifanos-neway/event-space-server/src/types"
	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	var signUpInput types.SignUpInput
	if err := c.BindJSON(&signUpInput); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid_Sign_Up_Input"})
		return
	}
	response := repos.SignupRepo(signUpInput)
	c.IndentedJSON(response.Code, gin.H{"message": response.Message})
}
