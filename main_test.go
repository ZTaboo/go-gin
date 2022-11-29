package main

import (
	"fmt"
	"testing"
	"zm/lib/z"
)

// jwt测试
func TestJwt(t *testing.T) {
	jwt, err := z.GenJwt()
	if err != nil {
		z.Log.Error("生成jwt错误：", err)
	}
	fmt.Println(jwt)
	status := z.CheckJwt(jwt)
	if status {
		z.Log.Info("is ok")
	} else {
		z.Log.Info("error")
	}
}
