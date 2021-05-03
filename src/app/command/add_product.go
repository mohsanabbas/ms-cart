package command

import (
	"context"
	"fmt"

	"github.com/mohsanabbas/ms-cart/src/domain/cart"
)



type AddProductHandler struct {
	cartRepo           cart.Repository
}

func NewAddProductHandler(cartRepo cart.Repository) AddProductHandler {
	if cartRepo == nil {
		panic("nil cartRepo")
	}
	return AddProductHandler{cartRepo:cartRepo}
}
func (h AddProductHandler) Handle(ctx context.Context,id string, cmd AddProduct) (err error) {
	defer func() {
		fmt.Printf("Error : %v/n" , err)
	}()

	p:=cart.Product{
		ID:   cmd.ID,
		Type: cmd.Type,
		Data: cmd.Data,
	}


	return h.cartRepo.AddProduct(
		ctx,
		id,
		&p,
		func(ctx context.Context, ct *cart.Cart) (*cart.Cart, error) {
			return ct, nil
		},
	)

}
