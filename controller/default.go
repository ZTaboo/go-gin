package controller

import (
	"github.com/gin-gonic/gin"
	"zm/lib/z"
)

func Hello(c *gin.Context) {
	z.HttpOk(c, z.H{
		Message: "hello world",
		Data:    nil,
	})
}
