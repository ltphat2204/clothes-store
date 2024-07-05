package main

import (
	"github.com/ltphat2204/clothes-store/pkg/server"
)

func main() {
	app := server.CreateServer()
	app.Run(":8080")
}