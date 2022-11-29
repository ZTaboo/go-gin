package z

import (
	"fmt"
	"testing"
)

func TestBcryptEncryption(t *testing.T) {
	str := BcryptEncryption("zero")
	fmt.Println(str)
	err := BcryptCheck(str, "zero2")
	if err != nil {
		fmt.Println("密码错误")
	} else {
		fmt.Println("密码正确")
	}
}
