package public

import (
	"github.com/PRSV-Hackathon-2025/go-backend/internal/wire"
	"github.com/gin-gonic/gin"
)

type HelloRouter struct {}

func (hr *HelloRouter) InitHelloRouter(r *gin.RouterGroup) {
	helloController, _ := wire.InitializeHelloController();
	
	helloRouter := r.Group("/hello")
	{
		helloRouter.GET("/", helloController.Hello)
	}
}