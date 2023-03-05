package repos

import (
	"log"

	"github.com/estifanos-neway/event-space-server/src/commons"
)

func SignOutRepo(refreshToken string) (int, string) {
	if err := DeleteRefreshToken(refreshToken); err != nil {
		log.Println("DeleteRefreshToken", err)
		return 500, commons.InternalError
	}
	return 200, commons.Ok
}
