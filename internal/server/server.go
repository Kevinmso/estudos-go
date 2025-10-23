package server

import (
	"log"

	"net/http"

	"github.com/kevinmso/estudos-go/internal/database"
	"github.com/kevinmso/estudos-go/internal/models"
	"github.com/labstack/echo/v4"
)

type Server interface {
	Start() error
	Readiness(ctx echo.Context) error
	Liveness(ctx echo.Context) error

	GetCustomersByEmail(ctx echo.Context) error
	GetAllVendors(ctx echo.Context) error
	GetProductsByVendor(ctx echo.Context) error
	GetAllServices(ctx echo.Context) error

	AddCustomer(ctx echo.Context) error
	AddVendor(ctx echo.Context) error
	AddProduct(ctx echo.Context) error
	AddService(ctx echo.Context) error
}

type EchoServer struct {
	echo *echo.Echo
	DB   database.DatabaseClient
}

func NewEchoServer(db database.DatabaseClient) Server {
	server := &EchoServer{
		echo: echo.New(),
		DB:   db,
	}
	server.registerRoutes()
	return server
}

func (s *EchoServer) Start() error {
	if err := s.echo.Start(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("could not start echo server: %v", err)
		return err
	}
	return nil
}

func (s *EchoServer) registerRoutes() {
	s.echo.GET("/readiness", s.Readiness)
	s.echo.GET("/liveness", s.Liveness)

	customerGroup := s.echo.Group("/customers")
	customerGroup.GET("", s.GetCustomersByEmail)
	customerGroup.POST("", s.AddCustomer)

	vendorGroup := s.echo.Group("/vendors")
	vendorGroup.GET("", s.GetAllVendors)
	vendorGroup.POST("", s.AddVendor)

	productGroup := s.echo.Group("/products")
	productGroup.GET("", s.GetProductsByVendor)
	productGroup.POST("", s.AddProduct)

	serviceGroup := s.echo.Group("/services")
	serviceGroup.GET("", s.GetAllServices)
	serviceGroup.POST("", s.AddService)
}

func (s *EchoServer) Readiness(ctx echo.Context) error {
	ready := s.DB.Ready()
	if ready {
		return ctx.JSON(http.StatusOK, models.Health{Status: "ready"})
	}
	return ctx.JSON(http.StatusServiceUnavailable, models.Health{Status: "unavailable"})
}

func (s *EchoServer) Liveness(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.Health{Status: "ok"})
}
