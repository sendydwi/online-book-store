package cart

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sendydwi/online-book-store/api"
	apicart "github.com/sendydwi/online-book-store/api/cart"
	"github.com/sendydwi/online-book-store/commons/utils"
	"github.com/sendydwi/online-book-store/services/product"
	"gorm.io/gorm"
)

type CartHandler struct {
	Svc Service
}

func NewRestHandler(db *gorm.DB) *CartHandler {
	return &CartHandler{
		Svc: Service{
			Repo: &CartRepository{DB: db},
			ProductSvc: &product.Service{
				Repo: &product.ProductRepository{
					DB: db,
				},
			},
		},
	}
}

func (c *CartHandler) RegisterHandler(g *gin.RouterGroup) {
	rGroup := g.Group("v1/carts")
	rGroup.POST("/", utils.CheckAuth, c.UpdateCartItem)
	rGroup.GET("/", utils.CheckAuth, c.GetCart)
}

func (c *CartHandler) UpdateCartItem(ctx *gin.Context) {
	var request apicart.CartUpdateRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		log.Printf("[update_cart_item][error] failed to read request %s", err.Error())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId := ctx.GetString("userId")
	err = c.Svc.UpdateCartItem(request, userId)
	if err != nil {
		log.Printf("[update_cart_item][error] failed to update cart item %s", err.Error())
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, api.GenericResponse{
		Status:  http.StatusOK,
		Message: "success",
	})
}

func (c *CartHandler) GetCart(ctx *gin.Context) {
	userId := ctx.GetString("userId")
	cartItemResponse, err := c.Svc.GetCartItem(userId)
	if err != nil {
		log.Printf("[update_cart_item][error] failed to update cart item %s", err.Error())
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, api.GenericResponseWithData{
		Status:  http.StatusOK,
		Message: "success",
		Data:    cartItemResponse,
	})
}
