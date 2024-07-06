package routes

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ltphat2204/clothes-store/pkg/common"
)

func createUploadRoutes(r *gin.RouterGroup) {
	r.POST("/upload", func(c *gin.Context) {
		log.Println("New file uploading")

		file, _ := c.FormFile("file")
		if file == nil {
			c.JSON(http.StatusBadRequest, common.NewSimpleError("No file uploaded"))
			return
		}

		now := time.Now().Format("20060102150405")
		dst := "assets/" + now + "_" + file.Filename

		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusInternalServerError, common.NewError(
				"Can not upload file",
				err.Error(),
			))
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "File uploaded successfully",
			"url": dst,
		})
	})
}