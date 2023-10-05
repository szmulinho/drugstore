package main

import (
	"github.com/szmulinho/drugstore/internal/database"
)

func main() {
	database.Connect()

	Run()
}
