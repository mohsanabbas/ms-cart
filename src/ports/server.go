package ports

import "github.com/labstack/echo/v4"

type HttpServerInterface interface {
	CreateCart(echo.Context) error
	AddProduct(echo.Context) error
	GetCart(echo.Context) error
}
