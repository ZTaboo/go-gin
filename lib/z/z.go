// Package z 个人封装部分
package z

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type H struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type CustomH struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func HttpOk(c *gin.Context, h H) {

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": h.Message,
		"data":    h.Data,
	})
}

func Error404(c *gin.Context, h H) {
	c.JSON(http.StatusNotFound, gin.H{
		"code":    404,
		"message": h.Message,
		"data":    h.Data,
	})
}

func Error400(c *gin.Context, h H) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    400,
		"message": h.Message,
		"data":    h.Data,
	})
}
