package handlers

import (
	"net/http"

	"github.com/estifanos-neway/event-space-server/src/commons"
	"github.com/estifanos-neway/event-space-server/src/repos"
	"github.com/estifanos-neway/event-space-server/src/types"
	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	var signUpInput types.SignUpInput
	if err := c.BindJSON(&signUpInput); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": commons.Invalid_Input})
		return
	}
	code, message := repos.SignupRepo(signUpInput)
	c.IndentedJSON(code, gin.H{"message": message})
}
