package web

import (
	"context"
	"ddd/internal/shopping_cart/application/checkout"
	"ddd/internal/shopping_cart/application/open_save"
	"ddd/internal/shopping_cart/application/send"
	"ddd/internal/shopping_cart/web/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShoppingCartsWebHandlers interface {
	Open(ctx *gin.Context)
	Save(ctx *gin.Context)
	Send(ctx *gin.Context)
	Checkout(ctx *gin.Context)
}

type shoppingCartsWebHandlersImpl struct {
	openAndSave          open_save.OpenAndSave
	sendShoppingCart     send.SendingShoppingCart
	checkoutShoppingCart checkout.CheckoutShoppingCart
}

func (handlers *shoppingCartsWebHandlersImpl) Open(ctx *gin.Context) {
	command := open_save.OpenCommand{}
	shoppingCart := handlers.openAndSave.Open(command)
	ctx.JSON(http.StatusOK, response.NewShoppingCartResponse(shoppingCart))
}

func (handlers *shoppingCartsWebHandlersImpl) Save(context context.Context) {}

func (handlers *shoppingCartsWebHandlersImpl) Send(context context.Context) {}

func (handlers *shoppingCartsWebHandlersImpl) Checkout(context context.Context) {}
