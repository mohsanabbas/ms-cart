package command

import (
	"context"
	"fmt"

	"github.com/mohsanabbas/ms-cart/src/domain/cart"
)



type CreatCartHandler struct {
	cartRepo           cart.Repository
}

func NewCreatCartHandler(cartRepo cart.Repository) CreatCartHandler {
	if cartRepo == nil {
		panic("nil cartRepo")
	}
	return CreatCartHandler{cartRepo:cartRepo}
}
func (h CreatCartHandler) Handle(ctx context.Context, cmd CreatCart) (err error) {
	defer func() {
		fmt.Printf("/n/n->CreatCart HERE----->,%v/n,%v/n" ,cmd.User, err)
	}()

	ct:=cart.Cart{
		ID:           cmd.ID,
		Expire:       cmd.Expire,
		Items:        cmd.Items,
		BusinessUnit: cmd.BusinessUnit,
		Agentsign:    cmd.Agentsign,
		User:         cmd.User,
	}


	if err := h.cartRepo.CreatCart(ctx, &ct); err != nil {
		return err
	}

	return nil
}
