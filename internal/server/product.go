package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *EchoServer) GetProductsByVendor(ctx echo.Context) error {
	vendor := ctx.QueryParam("vendor")

	products, err := s.DB.GetProductsByVendor(ctx.Request().Context(), vendor)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, products)
}
