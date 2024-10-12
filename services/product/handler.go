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
	rGroup := g.Group("v1/books")
	rGroup.GET("/:id", p.GetProductDetail)
	rGroup.GET("/", p.GetProductList)
}

func (p *ProductHandler) GetProductDetail(c *gin.Context) {
	bookId := c.Param("id")
	if bookId == "" {
		c.AbortWithError(http.StatusBadRequest, errors.New("id not found"))
	}

	bookDetail, err := p.Svc.GetBookById(bookId)
	//TODO fix error response
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("size is not a number"))
	}

	c.JSON(http.StatusOK, api.GenericResponseWithData{
		Status:  http.StatusOK,
		Message: "success",
		Data:    bookDetail,
	})
}

func (p *ProductHandler) GetProductList(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("page is not a number"))
	}
	size, err := strconv.Atoi(c.DefaultQuery("size", "0"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("size is not a number"))
	}

	bookList, err := p.Svc.GetBookList(page, size)
	//TODO fix error response
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("size is not a number"))
	}

	c.JSON(http.StatusOK, api.GenericResponseWithData{
		Status:  http.StatusOK,
		Message: "success",
		Data:    bookList,
	})
}