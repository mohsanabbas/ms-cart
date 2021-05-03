package query

import (
	"context"

	"github.com/mohsanabbas/ms-cart/src/domain/cart"
)

type GetCartReadModel interface {
	GetCart(ctx context.Context, id string) (*cart.Cart, error)
}

type GetCartHandler struct {
	cartRepo           cart.Repository
}

func NewGetCartHandler(cartRepo cart.Repository) GetCartHandler {
	if cartRepo == nil {
		panic("nil cartRepo")
	}

	return GetCartHandler{cartRepo:cartRepo}
}


func (h GetCartHandler) Handle(ctx context.Context, id string) (*cart.Cart,  error) {

	cart, err := h.cartRepo.GetCart(ctx,id)
	if err != nil {
		return nil, err
	}

	return cart, nil
}
