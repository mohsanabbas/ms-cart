package command

import (
	"github.com/mohsanabbas/ms-cart/src/domain/cart"
)

// type CreatCart struct {
// 	ID  primitive.ObjectID
// 	Items []cart.Product
// 	BusinessUnit string
// 	Expire int64
// 	User string
// 	Agentsign string
// }
type CreatCart cart.Cart

// type AddProduct struct{
// 	CartID string
// 	ID   primitive.ObjectID
// 	Type string
// 	Data map[string]interface{}
// }
type AddProduct cart.Product
