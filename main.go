package main

import (
	"log"

	"github.com/kevinmso/estudos-go/internal/database"
	"github.com/kevinmso/estudos-go/internal/server"
)

func main() {
	db, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatal("failed to connect to database, %s", err)
	}
	srv := server.NewEchoServer(db)
	if err := srv.Start(); err != nil {
		log.Fatal("failed to start server, %s", err)
	}
}
