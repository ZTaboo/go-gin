package router

import (
	"github.com/gin-gonic/gin"
	"zm/controller"
)

func hello(r *gin.RouterGroup) {
	v1 := controller.NewController()
	r.GET("/get", v1.Hello)
}
