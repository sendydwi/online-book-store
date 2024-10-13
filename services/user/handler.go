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

func (u *UserHandler) RegisterHandler(g *gin.RouterGroup) {
	rGroup := g.Group("v1/users")
	rGroup.POST("/register", u.RegisterUser)
	rGroup.POST("/login", u.LoginUser)
}

func (u *UserHandler) RegisterUser(ctx *gin.Context) {
	var request apiuser.CreateUserRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		log.Printf("[register_user][error] failed to read request %s", err.Error())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = u.Svc.RegisterUser(request.Email, request.Password)
	if err != nil {
		log.Printf("[register_user][error] failed to register user %s", err.Error())
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, api.GenericResponse{
		Status:  http.StatusOK,
		Message: "success",
	})
}

func (u *UserHandler) LoginUser(ctx *gin.Context) {
	var request apiuser.CreateUserRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		log.Printf("[register_user][error] failed to read request %s", err.Error())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := u.Svc.Login(request.Email, request.Password)
	if err != nil {
		log.Printf("[register_user][error] failed to read request %s", err.Error())
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, apiuser.LoginResponse{
		Token: token,
	})
}
