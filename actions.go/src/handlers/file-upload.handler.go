package handlers

import (
	"net/http"

	"github.com/estifanos-neway/event-space-server/src/commons"
	"github.com/estifanos-neway/event-space-server/src/repos"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func FileUploadHandler(c *gin.Context) {
	var body struct {
		Base64Str string `json:"base64Str"`
		Category  string `json:"category"`
		FileName  string `json:"fileName"`
		Extension string `json:"extension"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": commons.InvalidInput})
		return
	}
	const rootImagePath = "static/images/"
	var toPath string
	switch body.Category {
	case "event-image":
		toPath = rootImagePath + "events/"
	case "avatar-image":
		toPath = rootImagePath + "avatars/"
	default:
		toPath = rootImagePath + "others/"
	}
	fileName := body.FileName
	if fileName == "" {
		fileName = uuid.New().String() + body.Extension
	}
	toPath += fileName
	code, message := repos.FileUploadRepo(body.Base64Str, toPath)
	if code != 200 {
		c.IndentedJSON(code, gin.H{"message": message})
		return
	}
	c.IndentedJSON(code, gin.H{"message": message, "filePath": toPath})

}
