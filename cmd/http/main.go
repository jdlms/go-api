package main

import (
	"go-api/internal/auth"
	"go-api/internal/server"
)

func main() {

	auth.NewAuth()

	server.Startup()

}
