package initialize

import (
	"github.com/PRSV-Hackathon-2025/go-backend/global"
	"github.com/PRSV-Hackathon-2025/go-backend/internal/router"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine

	if global.Setting.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	r.Use() // logging
	r.Use() // cross
	r.Use() // limit rate

	publicRouter := router.RouterGroupApp.Public

	MainGroup := r.Group("/api/v1")
	{
		// MainGroup.GET("/health-check")
	}
	{
		publicRouter.InitHelloRouter(MainGroup)
	}

	return r
}