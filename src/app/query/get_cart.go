package query

import (
	"context"

	"github.com/mohsanabbas/ms-cart/src/domain/cart"
)

type GetCartReadModel interface {
	Handle(context.Context, string) (*cart.Cart, error)
}

type getCartHandler struct {
	cartRepo cart.Repository
}

func NewGetCartHandler(cartRepo cart.Repository) GetCartReadModel {
	if cartRepo == nil {
		panic("nil cartRepo")
	}

	return &getCartHandler{cartRepo: cartRepo}
}

func (h *getCartHandler) Handle(ctx context.Context, id string) (*cart.Cart, error) {

	cart, err := h.cartRepo.GetCart(ctx, id)
	if err != nil {
		return nil, err
	}

	return cart, nil
}
