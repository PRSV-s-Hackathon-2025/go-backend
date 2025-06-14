package router

import (
	"github.com/gin-gonic/gin"

)

func NewRouter() *gin.Engine {
	r := gin.Default()
	// r.Use(middleware.AuthMiddleware())

	return r
}
