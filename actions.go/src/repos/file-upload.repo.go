package repos

import (
	"log"

	"github.com/estifanos-neway/event-space-server/src/commons"
)

func FileUploadRepo(base64Str, toPath string) (int, string) {
	if err := commons.SaveFileFromBinary(base64Str, toPath); err != nil {
		log.Println("commons.SaveFileFromBinary", err)
		return 500, commons.InternalError
	}
	return 200, commons.Ok
}
