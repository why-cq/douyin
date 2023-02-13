package test

import (
	"douyin/utils"
	"fmt"
	"github.com/golang-jwt/jwt"
	"testing"
	"time"
)

// 用于签名的字符串  密钥，不能泄露
var mySigningKey = []byte("liwenzhou.com")

// GenRegisteredClaims 使用默认声明创建jwt
func GenRegisteredClaims() (string, error) {
	// 创建 Claims

	claims := &jwt.MapClaims{
		"ExpiresAt": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(), // 过期时间
		"Issuer":    "qimi",                                                // 签发人
	}
	// 生成token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成签名字符串
	return token.SignedString(mySigningKey)
}

// ParseRegisteredClaims 解析jwt
func ValidateRegisteredClaims(tokenString string) bool {
	// 解析token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil { // 解析token失败
		return false
	}
	return token.Valid
}

func TestGenRegisteredClaims(t *testing.T) {
	claims, err := GenRegisteredClaims()
	fmt.Println(len(claims))
	fmt.Println(err)

}

func TestValidateRegisteredClaims(t *testing.T) {
	claims, _ := GenRegisteredClaims()
	s := ValidateRegisteredClaims(claims)
	fmt.Println(s)
}

func TestGenerateToken(t *testing.T) {
	token1, err := utils.GenerateToken("why", "123456")
	token2, err := utils.GenerateToken("why", "123456")
	fmt.Println(token1.Values == token2.Values)
	fmt.Println(err)

}
