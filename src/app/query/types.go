package query

import (
	"github.com/mohsanabbas/ms-cart/src/domain/cart"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cart struct {
	ID           primitive.ObjectID
	Items        []cart.Product
	BusinessUnit string
	Expire       int64
	User         string
	Agentsign    string
}
