package handlers

import (
	"net/http"

	"github.com/estifanos-neway/event-space-server/src/commons"
	"github.com/estifanos-neway/event-space-server/src/repos"
	"github.com/gin-gonic/gin"
)

func VerifySignUpHandler(c *gin.Context) {
	var body struct{ VerificationToken string }
	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": commons.InvalidInput})
		return
	}
	if code, userLogin, message := repos.VerifySignupRepo(body.VerificationToken); userLogin == nil {
		c.IndentedJSON(code, gin.H{"message": message})
	} else {
		c.IndentedJSON(code, gin.H{"message": message, "userLogIn": *userLogin})
	}
}
