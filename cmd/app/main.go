package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/ltphat2204/clothes-store/internal/app"
)

func main() {
	app := server.CreateServer()
	app.Run(":8080")
}