//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"ddd/internal/shared"
	shoppingCart "ddd/internal/shopping_cart"
	shoppingCartWeb "ddd/internal/shopping_cart/web"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type WebServer struct {
	engine                  *gin.Engine
	shoppingCartWebHandlers shoppingCartWeb.ShoppingCartsWebHandlers
}

func (server *WebServer) Run(port string) {
	server.engine = gin.Default()
	server.engine.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello")
	})
	api := server.engine.Group("api")
	server.setupShoppingCartRoutes(api)
	server.engine.Run(port)
}

func (server *WebServer) setupShoppingCartRoutes(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	v1.POST("/shoppingcart", server.shoppingCartWebHandlers.Open)
	v1.PUT("/shoppingcart", server.shoppingCartWebHandlers.Save)
	v1.POST("/shoppingcart/send", server.shoppingCartWebHandlers.Send)
	v1.POST("/shoppingcart/checkout", server.shoppingCartWebHandlers.Checkout)
}

func buildWebServer(context context.Context) (WebServer, error) {
	wire.Build(shared.InfraSet, shared.EventBusSet, shoppingCart.ShoppCartWeb, wire.Struct(new(WebServer), "shoppingCartWebHandlers"))
	return WebServer{}, nil
}
