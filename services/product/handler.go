package product

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sendydwi/online-book-store/api"
	"gorm.io/gorm"
)

type ProductHandler struct {
	Svc ProductServiceInterface
}

func NewRestHandler(db *gorm.DB) *ProductHandler {
	return &ProductHandler{
		Svc: &Service{
			Repo: &ProductRepository{DB: db},
		},
	}
}

func (p *ProductHandler) RegisterHandler(g *gin.RouterGroup) {
	rGroup := g.Group("v1/products")
	rGroup.GET("/:id", p.GetProductDetail)
	rGroup.GET("/", p.GetProductList)
}

func (p *ProductHandler) GetProductDetail(ctx *gin.Context) {
	productId := ctx.Param("id")
	if productId == "" {
		ctx.JSON(http.StatusBadRequest, api.GenericResponse{
			Status:  http.StatusBadRequest,
			Message: "id param not found",
		})
		return
	}

	productIdValue, err := strconv.Atoi(productId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, api.GenericResponse{
			Status:  http.StatusBadRequest,
			Message: "id is not a number",
		})
		return
	}

	productDetail, err := p.Svc.GetProductById(productIdValue)
	//TODO fix error response
	if err != nil {
		log.Printf("[get_product_detail][error] failed to get product with id %v with error: %s", productIdValue, err.Error())
		switch {
		case errors.Is(err, ErrProductNotFound):
			ctx.JSON(http.StatusNotFound, api.GenericResponse{
				Status:  http.StatusNotFound,
				Message: fmt.Sprintf("product with id %v not found", productIdValue),
			})
		default:
			ctx.JSON(http.StatusInternalServerError, api.GenericResponse{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, api.GenericResponseWithData{
		Status:  http.StatusOK,
		Message: "success",
		Data:    productDetail,
	})
}

func (p *ProductHandler) GetProductList(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, api.GenericResponse{
			Status:  http.StatusBadRequest,
			Message: "page is not a number",
		})
		return
	}
	size, err := strconv.Atoi(ctx.DefaultQuery("size", "10"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, api.GenericResponse{
			Status:  http.StatusBadRequest,
			Message: "size is not a number",
		})
		return
	}

	bookList, err := p.Svc.GetProductList(page, size)
	//TODO fix error response
	if err != nil {
		log.Printf("[get_product_detail][error] failed to get product list with error: %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, api.GenericResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, api.GenericResponseWithData{
		Status:  http.StatusOK,
		Message: "success",
		Data:    bookList,
	})
}
