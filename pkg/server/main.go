package server

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ltphat2204/clothes-store/pkg/routes"
)

func init() {
	godotenv.Load(".env.local")
	godotenv.Load(".env.development")
	godotenv.Load(".env.production")
	godotenv.Load(".env")
}

type server struct {
	router *gin.Engine
}
  
func CreateServer() *server {
	app := gin.Default()

	app.StaticFS("/assets", http.Dir("assets"))

	routes.CreateRoutesV1(app)

	return &server{
		router: app,
	}
}

func (s *server) Run(port string) {
	s.router.Run(port)
}