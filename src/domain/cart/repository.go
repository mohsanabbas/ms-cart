package cart

import (
	"context"
	"fmt"
)

type NotFoundError struct {
	id string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("Cart '%s' not found", e.id)
}

type Repository interface {
	CreatCart(ctx context.Context, ct *Cart) error
	GetCart(ctx context.Context, id string) (*Cart, error)
	AddProduct(
		ctx context.Context,
		id string,
		ct *Product,
		updateFn func(ctx context.Context, ct *Cart) (*Cart, error),
	) error
}
