package handlers

import (
	"net/http"
	"strings"

	"github.com/estifanos-neway/event-space-server/src/commons"
	"github.com/estifanos-neway/event-space-server/src/repos"
	"github.com/estifanos-neway/event-space-server/src/types"
	"github.com/gin-gonic/gin"
)

func SignInHandler(c *gin.Context) {
	var signInInput types.SignInInput
	if err := c.BindJSON(&signInInput); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": commons.InvalidInput})
		return
	}
	signInInput.Email = strings.ToLower(signInInput.Email)
	if code, userLogin, message := repos.SignInRepo(signInInput); userLogin == nil {
		c.IndentedJSON(code, gin.H{"message": message})
	} else {
		c.IndentedJSON(code, gin.H{"message": message, "userLogIn": *userLogin})
	}
}
