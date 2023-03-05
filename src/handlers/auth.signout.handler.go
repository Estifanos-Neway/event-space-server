package handlers

import (
	"net/http"

	"github.com/estifanos-neway/event-space-server/src/commons"
	"github.com/estifanos-neway/event-space-server/src/repos"
	"github.com/gin-gonic/gin"
)

func SignOutHandler(c *gin.Context) {
	var body struct{ RefreshToken string }
	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": commons.InvalidInput})
		return
	}
	code, message := repos.SignOutRepo(body.RefreshToken)
	c.IndentedJSON(code, gin.H{"message": message})
}
