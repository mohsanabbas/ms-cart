package main

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/mohsanabbas/ms-cart/src/ports"
	"github.com/mohsanabbas/ms-cart/src/service"
	"github.com/spf13/viper"
)

var (
	e = echo.New()
)

func main() {
	// e.Use(middleware.Recover())
	// middlewares
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10,
		LogLevel:  log.ERROR,
	}))
	e.Use(middleware.Timeout())

	ctx := context.Background()

	application := service.NewApplication(ctx)

	ctHandler := ports.NewHttpServer(e, application)

	e.POST("/cart", ctHandler.CreateCart)
	e.GET("/cart/:id", ctHandler.GetCart)
	e.PATCH("/cart/:id", ctHandler.AddProduct)

	if err := e.Start(fmt.Sprintf("%v", viper.Get("SERVER_ADDRESS"))); err != nil {
		fmt.Println(err)
	}

}
