package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sendydwi/online-book-store/api"
	apiuser "github.com/sendydwi/online-book-store/api/user"
	"gorm.io/gorm"
)

type UserHandler struct {
	Svc Service
}

func NewRestHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		Svc: Service{
			Repo: UserRepository{DB: db},
		},
	}
}

func (u *UserHandler) RegisterHandler(r *gin.Engine) {
	rGroup := r.Group("api/v1/user")
	rGroup.POST("/register", u.RegisterUser)
	rGroup.POST("/login", u.LoginUser)
	rGroup.POST("/logout", u.LogoutUser)
}

func (u *UserHandler) RegisterUser(c *gin.Context) {
	var request apiuser.CreateUserRequest
	err := c.BindJSON(&request)
	if err != nil {
		log.Printf("[register_user][error] failed to read request %s", err.Error())
	}

	err = u.Svc.RegisterUser(request.Email, request.Password)
	if err != nil {
		log.Printf("[register_user][error] failed to register user %s", err.Error())
	}

	c.JSON(http.StatusOK, api.GenericResponse{
		Status:  http.StatusOK,
		Message: "success",
	})
}

func (u *UserHandler) LoginUser(c *gin.Context) {
	//TODO
}

func (u *UserHandler) LogoutUser(c *gin.Context) {
	//TODO
}
