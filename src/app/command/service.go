package command

import (
	"context"

	"github.com/mohsanabbas/ms-cart/src/domain/cart"
)
type CartService interface {
	CreatCart(ctx context.Context, cart cart.Cart) error
	AddProduct(ctx context.Context,id string, cart cart.Product) error
	// RemoveProduct(ctx context.Context, id string) error
}
