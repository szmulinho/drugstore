package main

import (
	"github.com/szmulinho/drugstore/cmd/server"
	"github.com/szmulinho/drugstore/database"
)

func main() {
	database.Connect()

	server.Run()
}
