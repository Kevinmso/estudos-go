package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *EchoServer) GetProductsByVendor(ctx echo.Context) error {
	vendor := ctx.Param("vendor")

	products, err := s.DB.GetProductsByVendor(ctx.Request().Context(), vendor)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, products)
}
