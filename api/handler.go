package api

import "github.com/gin-gonic/gin"

type Handler interface {
	RegisterHandler(r *gin.Engine)
}
