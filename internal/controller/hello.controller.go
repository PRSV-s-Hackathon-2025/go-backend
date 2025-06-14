package controller

import (
	"net/http"

	"github.com/PRSV-Hackathon-2025/go-backend/pkg/response"
	"github.com/gin-gonic/gin"
)

type HelloController struct {
}

func NewHelloController() *HelloController {
	return &HelloController{}
}

func (hc *HelloController) Hello(c *gin.Context) {
	response.SuccessResponse(c, http.StatusOK, "hello")
}
