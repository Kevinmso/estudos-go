package server

import (
	"log"

	"net/http"

	"github.com/kevinmso/estudos-go/internal/database"
	"github.com/labstack/echo/v4"
)

type Server interface {
	Start() error
}

type EchoServer struct {
	echo *echo.Echo
	DB   *database.DatabaseClient
}

func newEchoServer(db *database.DatabaseClient) Server {
	server := &EchoServer{
		echo: echo.New(),
		DB:   db,
	}
	server.registerRoutes()
	return server
}

func (e *EchoServer) Start() error {
	if err := e.echo.Start(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("could not start echo server: %v", err)
		return err
	}
	return nil
}

func (e *EchoServer) registerRoutes() {

}
