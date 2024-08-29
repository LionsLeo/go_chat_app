package main

import (
	"chat_server/db"
	"log"
)

func main() {
	_, err := db.NewDatabase()
	if err != nil {
		log.Fatal("could not initialize database connection: %s", err)
	}
}
