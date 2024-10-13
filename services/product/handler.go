package product

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sendydwi/online-book-store/api"
	"gorm.io/gorm"
)

type ProductHandler struct {
	Svc Service
}

func NewRestHandler(db *gorm.DB) *ProductHandler {
	return &ProductHandler{
		Svc: Service{
			Repo: ProductRepository{DB: db},
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
		ctx.AbortWithError(http.StatusBadRequest, errors.New("id not found"))
		return
	}

	productIdValue, err := strconv.Atoi(productId)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("id is not a number"))
		return
	}

	productDetail, err := p.Svc.GetProductById(productIdValue)
	//TODO fix error response
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("id is not a number"))
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
		ctx.AbortWithError(http.StatusBadRequest, errors.New("page is not a number"))
		return
	}
	size, err := strconv.Atoi(ctx.DefaultQuery("size", "10"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("size is not a number"))
		return
	}

	bookList, err := p.Svc.GetProductList(page, size)
	//TODO fix error response
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("size is not a number"))
		return
	}

	ctx.JSON(http.StatusOK, api.GenericResponseWithData{
		Status:  http.StatusOK,
		Message: "success",
		Data:    bookList,
	})
}
