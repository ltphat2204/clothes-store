package main

import (
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

func main() {
	app := gin.Default()
	routes.CreateRoutesV1(app)
	app.Run()
}