package repos

import (
	"log"

	"github.com/estifanos-neway/event-space-server/src/commons"
	"github.com/estifanos-neway/event-space-server/src/types"
)

func TicketSellRepo(newTicket types.Ticket) (int, string) {
	// get user email
	user, err := getUserById(uuid(newTicket.UserId))
	if err != nil {
		log.Println("getUserById", err)
		return 500, commons.InternalError
	}
	// create qr code
	qrUrl := newTicket.Id
	qrPath := "static/images/ticket-qr-codes/" + newTicket.Id + ".png"
	if err := commons.CreateQrCode(qrUrl, qrPath); err != nil {
		log.Println("CreateQrCode", err)
		return 500, commons.InternalError
	}
	// send email
	planeContent := "You have successfully bought a ticket."
	subject := "New Ticket"
	if err := commons.SendEmail(user.Email, &planeContent, nil, nil, &subject, &qrPath); err != nil {
		log.Println("SendEmail", err)
		return 500, commons.InternalError
	}
	// response
	return 200, commons.Ok
}
