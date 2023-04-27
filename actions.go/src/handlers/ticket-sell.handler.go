package handlers

import (
	"net/http"

	"github.com/estifanos-neway/event-space-server/src/commons"
	"github.com/estifanos-neway/event-space-server/src/repos"
	"github.com/gin-gonic/gin"
)

func TicketSellHandler(c *gin.Context) {
	event := newTicketEvent{}
	if err := c.BindJSON(&event); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": commons.InvalidInput})
		return
	}
	code, message := repos.TicketSellRepo(event.Event.Data.New)
	c.IndentedJSON(code, gin.H{"message": message})
}
