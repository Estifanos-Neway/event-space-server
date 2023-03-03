package handlers

import (
	"net/http"

	"github.com/estifanos-neway/event-space-server/src/commons"
	"github.com/estifanos-neway/event-space-server/src/repos"
	"github.com/gin-gonic/gin"
)

func RefreshHandler(c *gin.Context) {
	var body struct{ AccessToken string }
	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": commons.Invalid_Input})
		return
	}
	if code, userLogin, message := repos.RefreshRepo(body.AccessToken); userLogin == nil {
		c.IndentedJSON(code, gin.H{"message": message})
	} else {
		c.IndentedJSON(code, *userLogin)
	}
}