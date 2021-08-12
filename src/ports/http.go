package ports

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mohsanabbas/ms-cart/src/app"
	"github.com/mohsanabbas/ms-cart/src/app/command"
	"github.com/mohsanabbas/ms-cart/src/domain/cart"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type httpServer struct {
	app app.Application
}

func NewHttpServer(e *echo.Echo, application app.Application) HttpServerInterface {
	return &httpServer{
		app: application,
	}
}

func (h *httpServer) CreateCart(c echo.Context) (err error) {
	cmd := cart.Cart{
		ID: primitive.NewObjectID(),
	}
	if err := c.Bind(&cmd); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = h.app.Commands.CreatCart.Handle(c.Request().Context(), command.CreatCart(cmd))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	c.Response().WriteHeader(http.StatusCreated)
	return json.NewEncoder(c.Response()).Encode(cmd)
}

func (h *httpServer) AddProduct(c echo.Context) (err error) {
	cartID := c.Param("id")
	addProduct := cart.Product{
		ID: primitive.NewObjectID(),
	}
	if err := c.Bind(&addProduct); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err = h.app.Commands.AddProduct.Handle(c.Request().Context(), cartID, command.AddProduct(addProduct))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(addProduct)
}

func (h *httpServer) GetCart(c echo.Context) (err error) {

	cartID := c.Param("id")
	cart, err := h.app.Queries.GetCart.Handle(c.Request().Context(), cartID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, cart)

}
