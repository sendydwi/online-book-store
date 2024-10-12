package cart

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CartHandler struct {
	Svc Service
}

func NewRestHandler(db *gorm.DB) *CartHandler {
	return &CartHandler{
		Svc: Service{
			Repo: Repository{DB: db},
		},
	}
}

func (c *CartHandler) RegisterHandler(r *gin.Engine) {
	rGroup := r.Group("v1/carts")
	rGroup.POST("/update", c.UpdateCartItem)
	rGroup.GET("/", c.GetCart)
}

func (c *CartHandler) UpdateCartItem(ctx *gin.Context) {

}

func (c *CartHandler) GetCart(ctx *gin.Context) {

}
