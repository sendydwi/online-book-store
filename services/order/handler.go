package order

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sendydwi/online-book-store/api"
	apiorder "github.com/sendydwi/online-book-store/api/order"
	"github.com/sendydwi/online-book-store/commons/utils"
	"github.com/sendydwi/online-book-store/services/cart"
	"github.com/sendydwi/online-book-store/services/product"
	"gorm.io/gorm"
)

type OrderHandler struct {
	Svc Service
}

func NewRestHandler(db *gorm.DB) *OrderHandler {
	productSvc := product.Service{
		Repo: product.ProductRepository{
			DB: db,
		},
	}

	return &OrderHandler{
		Svc: Service{
			Repo: OrderRepository{DB: db},
			CartSvc: cart.Service{
				Repo: cart.CartRepository{
					DB: db,
				},
				ProductSvc: productSvc,
			},
			ProductSvc: productSvc,
		},
	}
}

func (o *OrderHandler) RegisterHandler(g *gin.RouterGroup) {
	rGroup := g.Group("v1/orders")
	rGroup.POST("/", utils.CheckAuth, o.CreateOrder)
	rGroup.GET("/:id", utils.CheckAuth, o.GetOrderDetail)
	rGroup.GET("/", utils.CheckAuth, o.GetOrderHistories)
}

func (o *OrderHandler) CreateOrder(ctx *gin.Context) {
	var request apiorder.CreateOrderRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		log.Printf("[create_order][error] failed to read request, error: %s", err.Error())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId := ctx.GetString("userId")
	err = o.Svc.CreateOrder(userId, request)
	if err != nil {
		log.Printf("[create_order][error] failed to create order, error: %s", err.Error())
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, api.GenericResponse{
		Status:  http.StatusCreated,
		Message: "created",
	})
}

func (o *OrderHandler) GetOrderDetail(ctx *gin.Context) {
	orderId := ctx.Param("id")
	if orderId == "" {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("id not found"))
		return
	}

	userId := ctx.GetString("userId")
	response, err := o.Svc.GetOrderDetail(orderId, userId)
	if err != nil {
		log.Printf("[get_order_history][error] failed to create order, error: %s", err.Error())
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, api.GenericResponseWithData{
		Status:  http.StatusOK,
		Message: "created",
		Data:    response,
	})
}

func (o *OrderHandler) GetOrderHistories(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "0"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("page is not a number"))
		return
	}
	size, err := strconv.Atoi(ctx.DefaultQuery("size", "0"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("size is not a number"))
		return
	}

	userId := ctx.GetString("userId")
	response, err := o.Svc.GetOrderHistories(userId, page, size)
	if err != nil {
		log.Printf("[get_order_history][error] failed to create order, error: %s", err.Error())
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, api.GenericResponseWithData{
		Status:  http.StatusOK,
		Message: "created",
		Data:    response,
	})
}
