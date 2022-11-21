package router

import (
	"github.com/gin-gonic/gin"
	"zm/middleware"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors()) //跨域中间件
	v1 := r.Group("/v1")
	hello(v1)
	return r
}
