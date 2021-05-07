package service

import (
	"context"

	"github.com/mohsanabbas/ms-cart/src/adapters"
	"github.com/mohsanabbas/ms-cart/src/app"
	"github.com/mohsanabbas/ms-cart/src/app/command"
	"github.com/mohsanabbas/ms-cart/src/app/query"
	"github.com/mohsanabbas/ms-cart/src/clients/mongo"
)

func NewApplication(ctx context.Context) app.Application {


	// config, err := config.GetConfig(".")
	// if err != nil {
	// 	fmt.Printf("Error connecting to database:%v", err)
	// }

	// GetSession gets mongo client
	session := mongo.GetSession()
	// defer session.Disconnect(ctx)


	// client := session.Database(config.DBName).Collection(config.Collection)
	cartRepository := adapters.NewCartMongoRepository(session)


	return app.Application{
		Commands: app.Commands{
			CreatCart:     command.NewCreatCartHandler(cartRepository),
			AddProduct:       command.NewAddProductHandler(cartRepository),
		},
		Queries: app.Queries{
			GetCart: query.NewGetCartHandler(cartRepository),
		},
	}
}
