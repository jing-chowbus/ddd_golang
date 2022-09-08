package web

import (
	"ddd/internal/shopping_cart/application/checkout"
	"ddd/internal/shopping_cart/application/open_save"
	"ddd/internal/shopping_cart/application/send"
	"ddd/internal/shopping_cart/domain"
	"ddd/internal/shopping_cart/web/request"
	"ddd/internal/shopping_cart/web/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ShoppingCartsWebHandlers interface {
	Open(ctx *gin.Context)
	Save(ctx *gin.Context)
	Send(ctx *gin.Context)
	Checkout(ctx *gin.Context)
	Find(ctx *gin.Context)
	FindByID(ctx *gin.Context)
}

type shoppingCartsWebHandlersImpl struct {
	openAndSave          open_save.OpenAndSave
	sendShoppingCart     send.SendingShoppingCart
	checkoutShoppingCart checkout.CheckoutShoppingCart
	shoppingCarts        domain.ShoppingCarts
}

func NewShopingCartsWebHanders(openAndSave open_save.OpenAndSave, sendShoppingCart send.SendingShoppingCart, checkoutShoppingCart checkout.CheckoutShoppingCart, shoppingCarts domain.ShoppingCarts) ShoppingCartsWebHandlers {
	return &shoppingCartsWebHandlersImpl{
		openAndSave,
		sendShoppingCart,
		checkoutShoppingCart,
		shoppingCarts,
	}
}

func (handlers *shoppingCartsWebHandlersImpl) Open(ctx *gin.Context) {
	command := open_save.OpenCommand{}
	shoppingCart := handlers.openAndSave.Open(command)
	ctx.JSON(http.StatusOK, response.NewShoppingCartResponse(shoppingCart))
}

func (handlers *shoppingCartsWebHandlersImpl) Save(ctx *gin.Context) {}

func (handlers *shoppingCartsWebHandlersImpl) Send(ctx *gin.Context) {}

func (handlers *shoppingCartsWebHandlersImpl) Checkout(ctx *gin.Context) {}

func (handlers *shoppingCartsWebHandlersImpl) FindByID(ctx *gin.Context) {
	shoppingCartID, err := uuid.Parse(ctx.Param("shopping_cart_id"))
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	shoppingCart, err := handlers.shoppingCarts.FindByID(shoppingCartID)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	ctx.JSON(http.StatusOK, response.NewShoppingCartResponse(shoppingCart))
}

func (handlers *shoppingCartsWebHandlersImpl) Find(ctx *gin.Context) {
	var queryRequest request.QueryShoppingCarts
	if err := ctx.ShouldBind(&queryRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, err)
	}
}
