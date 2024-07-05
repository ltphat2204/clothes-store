package routes

import (
	"net/http"
  
	"github.com/gin-gonic/gin"
)
  
func CreateRoutes(app *gin.Engine) {
	r := app.Group("/v1")
	{
		r.GET("/helloworld", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello world!",
			})
		})
	}
}