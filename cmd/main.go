package main

import (
	"github.com/szmulinho/drugstore/cmd/server"
	"github.com/szmulinho/drugstore/internal/database"
)

func main() {
	database.Connect()

	server.Run()
}
