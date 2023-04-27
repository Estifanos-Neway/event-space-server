package handlers

import (
	"net/http"
	"strings"

	"github.com/estifanos-neway/event-space-server/src/commons"
	"github.com/estifanos-neway/event-space-server/src/repos"
	"github.com/estifanos-neway/event-space-server/src/types"
	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	var signUpInput types.SignUpInput
	if err := c.BindJSON(&signUpInput); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": commons.InvalidInput})
		return
	}
	signUpInput.Email = strings.ToLower(signUpInput.Email)
	code, message := repos.SignUpRepo(signUpInput)
	c.IndentedJSON(code, gin.H{"message": message})
}
