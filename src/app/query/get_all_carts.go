package query

import (
	"context"
)

type AllCartsReadModel interface {
	Handle(ctx context.Context) ([]Cart, error)
}
type allCartsHandler struct {
	readModel AllCartsReadModel
}

func NewAllCartsHandler(readModel AllCartsReadModel) AllCartsReadModel {
	if readModel == nil {
		panic("nil readModel")
	}

	return &allCartsHandler{readModel: readModel}
}

func (h *allCartsHandler) Handle(ctx context.Context) (tr []Cart, err error) {
	return h.readModel.Handle(ctx)
}
