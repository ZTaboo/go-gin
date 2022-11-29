package z

import (
	"golang.org/x/crypto/bcrypt"
)

// BcryptEncryption bcrypt方式加密
func BcryptEncryption(str string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		Log.Error("加密错误：", err)
	}
	return string(hash)
}

// BcryptCheck 校验密码正确性
// pass为密码，str为需要被验证正确性的字符串
func BcryptCheck(pass, str string) error {
	err := bcrypt.CompareHashAndPassword([]byte(pass), []byte(str))
	if err != nil {
		Log.Error("校验错误:", err)
		return err
	}
	return nil
}
