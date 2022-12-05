package controller

import (
	"github.com/gin-gonic/gin"
	"zm/lib/z"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

// @BasePath /api/v1

// Hello
// PingExample godoc
// @Summary 测试 Summary
// @Schemes
// @Description 测试 Description
// @Tags v1
// @Accept json
// @Produce json
// @Success 200 {string} Hello
// @Router /v1/Hello [get]
func (r Controller) Hello(c *gin.Context) {
	z.HttpOk(c, z.H{
		Message: "hello world",
	})
}
