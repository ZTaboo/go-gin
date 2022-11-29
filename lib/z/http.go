// Package z 个人封装部分
package z

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type H struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Total   int    `json:"total,omitempty"`
}

func HttpOk(c *gin.Context, h H) {
	var data gin.H
	h.Code = 200
	marshal, _ := json.Marshal(h)
	err := json.Unmarshal(marshal, &data)
	if err != nil {
		Log.Error("解析错误参数", err)
	}
	c.JSON(http.StatusOK, data)
}

func Error404(c *gin.Context, h H) {
	var data gin.H
	h.Code = 404
	marshal, _ := json.Marshal(h)
	err := json.Unmarshal(marshal, &data)
	if err != nil {
		Log.Error("解析错误参数", err)
	}
	c.JSON(http.StatusNotFound, data)
}

func Error400(c *gin.Context, h H) {
	var data gin.H
	h.Code = 400
	marshal, _ := json.Marshal(h)
	err := json.Unmarshal(marshal, &data)
	if err != nil {
		Log.Error("解析错误参数", err)
	}
	c.JSON(http.StatusBadRequest, data)
}
