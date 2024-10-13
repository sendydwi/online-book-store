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

func (p *ProductHandler) RegisterHandler(g *gin.Engine) {
	rGroup := g.Group("v1/books")
	rGroup.GET("/:id", p.GetProductDetail)
	rGroup.GET("/", p.GetProductList)
}

func (p *ProductHandler) GetProductDetail(ctx *gin.Context) {
	bookId := ctx.Param("id")
	if bookId == "" {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("id not found"))
	}

	bookIdValue, err := strconv.Atoi(bookId)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("id is not a number"))
	}

	bookDetail, err := p.Svc.GetProductById(bookIdValue)
	//TODO fix error response
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("id is not a number"))
	}

	ctx.JSON(http.StatusOK, api.GenericResponseWithData{
		Status:  http.StatusOK,
		Message: "success",
		Data:    bookDetail,
	})
}

func (p *ProductHandler) GetProductList(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "0"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("page is not a number"))
	}
	size, err := strconv.Atoi(ctx.DefaultQuery("size", "0"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("size is not a number"))
	}

	bookList, err := p.Svc.GetProductList(page, size)
	//TODO fix error response
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("size is not a number"))
	}

	ctx.JSON(http.StatusOK, api.GenericResponseWithData{
		Status:  http.StatusOK,
		Message: "success",
		Data:    bookList,
	})
}
