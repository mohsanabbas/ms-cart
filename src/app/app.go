package app

import (
	"github.com/mohsanabbas/ms-cart/src/app/command"
	"github.com/mohsanabbas/ms-cart/src/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreatCart command.CreatCartHandler
	AddProduct command.AddProductHandler
}

type Queries struct {
	GetCart query.GetCartHandler
	AllCarts query.AllCartsHandler
}
