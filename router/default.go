package router

import (
	"github.com/gin-gonic/gin"
	"zm/controller"
)

func hello(v1 *gin.RouterGroup) {
	v1.GET("/get", controller.Hello)
}
