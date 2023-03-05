package handlers

import (
	"github.com/estifanos-neway/event-space-server/src/repos"
	"github.com/estifanos-neway/event-space-server/src/types"
)

type signInResponse struct {
	success   bool
	message   string
	userLogIn repos.UserLogin
}

type newTicketEvent struct {
	Event struct {
		Data struct {
			New types.Ticket `json:"New"`
		} `json:"data"`
	} `json:"event"`
}
