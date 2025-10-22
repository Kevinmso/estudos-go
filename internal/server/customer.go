package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *EchoServer) GetCustomersByEmail(ctx echo.Context) error {
	email := ctx.QueryParam("email")

	customers, err := s.DB.GetCustomersByEmail(ctx.Request().Context(), email)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, customers)
}
