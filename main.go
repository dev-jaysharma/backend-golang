package main

import (
	"log"

	"github.com/dev-jaysharma/backend-golang/db"
)

func main() {
	db.Connector()
	log.Println("🚀Connected to database by main.go file")
}