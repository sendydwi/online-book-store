package user

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sendydwi/online-book-store/api"
	apiuser "github.com/sendydwi/online-book-store/api/user"
	"gorm.io/gorm"
)

type UserHandler struct {
	Svc UserServiceInterface
}

func NewRestHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		Svc: &Service{
			Repo: &UserRepository{DB: db},
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

		switch {
		case errors.Is(err, ErrEmailAlreadyUsed):
			ctx.JSON(http.StatusUnprocessableEntity, api.GenericResponse{
				Status:  http.StatusUnprocessableEntity,
				Message: "email already used",
			})
		default:
			ctx.JSON(http.StatusInternalServerError, api.GenericResponse{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
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
		ctx.JSON(http.StatusBadRequest, api.GenericResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	token, err := u.Svc.Login(request.Email, request.Password)
	if err != nil {
		log.Printf("[register_user][error] failed to login with email %s : %s", request.Email, err.Error())

		switch {
		case errors.Is(err, ErrUserNotExist), errors.Is(err, ErrWrongPassword):
			ctx.JSON(http.StatusUnprocessableEntity, api.GenericResponse{
				Status:  http.StatusUnprocessableEntity,
				Message: "wrong email or password",
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
		Data: apiuser.LoginResponse{
			Token: token,
		},
	})
}
