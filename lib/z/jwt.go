package z

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
	"zm/config"
)

func GenJwt() (string, error) {
	// 创建 Claims
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 过期时间
		Issuer:    config.Info.Jwt.Issuer,                             // 签发人
	}
	// 生成token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成签名字符串
	return token.SignedString([]byte(config.Info.Jwt.SigningKey))
}

func CheckJwt(tokenString string) bool {
	// 解析token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Info.Jwt.SigningKey), nil
	})
	if err != nil { // 解析token失败
		Log.Error("parse token error：", err)
		return false
	}
	return token.Valid
}
