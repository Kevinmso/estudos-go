package server

import (
	"net/http"

	"github.com/kevinmso/estudos-go/internal/dberrors"
	"github.com/kevinmso/estudos-go/internal/models"
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

func (s *EchoServer) AddProduct(ctx echo.Context) error {
	product := new(models.Product)
	if err := ctx.Bind(product); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	product, err := s.DB.AddProduct(ctx.Request().Context(), product)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, product)
}
