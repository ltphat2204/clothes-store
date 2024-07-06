package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ltphat2204/clothes-store/internal/routes/v1"
)

type server struct {
	router *gin.Engine
}
  
func CreateServer() *server {
	app := gin.Default()

	app.StaticFS("/assets", http.Dir("assets"))

	routes.CreateRoutes(app)

	return &server{
		router: app,
	}
}

func (s *server) Run(port string) {
	s.router.Run(port)
}