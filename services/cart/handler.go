package cart

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sendydwi/online-book-store/api"
	apicart "github.com/sendydwi/online-book-store/api/cart"
	"gorm.io/gorm"
)

type CartHandler struct {
	Svc Service
}

func NewRestHandler(db *gorm.DB) *CartHandler {
	return &CartHandler{
		Svc: Service{
			Repo: CartRepository{DB: db},
		},
	}
}

func (c *CartHandler) RegisterHandler(r *gin.Engine) {
	rGroup := r.Group("v1/carts")
	rGroup.POST("/update", c.UpdateCartItem)
	rGroup.GET("/", c.GetCart)
}

func (c *CartHandler) UpdateCartItem(ctx *gin.Context) {
	var request apicart.CartUpdateRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		log.Printf("[update_cart_item][error] failed to read request %s", err.Error())
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	err = c.Svc.UpdateCartItem(request, "")
	if err != nil {
		log.Printf("[update_cart_item][error] failed to update cart item %s", err.Error())
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, api.GenericResponse{
		Status:  http.StatusOK,
		Message: "success",
	})
}

func (c *CartHandler) GetCart(ctx *gin.Context) {
	cartItemResponse, err := c.Svc.GetCartItem("")
	if err != nil {
		log.Printf("[update_cart_item][error] failed to update cart item %s", err.Error())
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, api.GenericResponseWithData{
		Status:  http.StatusOK,
		Message: "success",
		Data:    cartItemResponse,
	})
}
