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

func (r Controller) Hello(c *gin.Context) {
	z.HttpOk(c, z.H{
		Message: "hello world",
	})
}
