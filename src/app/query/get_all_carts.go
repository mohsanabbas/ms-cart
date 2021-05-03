package query

import (
	"context"
)

type AllCartsReadModel interface {
	AllCarts(ctx context.Context) ([]Cart, error)
}
type AllCartsHandler struct {
	readModel AllCartsReadModel
}

func NewAllCartsHandler(readModel AllCartsReadModel) AllCartsHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return AllCartsHandler{readModel: readModel}
}



func (h AllCartsHandler) Handle(ctx context.Context) (tr []Cart, err error) {
	return h.readModel.AllCarts(ctx)
}
