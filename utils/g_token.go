package utils

import (
	"github.com/golang-jwt/jwt"
	"time"
)

// 返回的token字段
type Token struct {
	Values string
}

// 用于签名的字符串  密钥，不能泄露
var mySigningKey = []byte("why")

func GenerateToken(username string, password string) (Token, error) {
	// 创建 Claims
	claims := &jwt.MapClaims{
		"username":  username,
		"password":  password,
		"ExpiresAt": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(), // 过期时间
		"Issuer":    "qimi",                                                // 签发人
	}
	// 生成token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成签名字符串
	s, err := token.SignedString(mySigningKey)
	if err != nil {
		return Token{Values: s}, err
	}
	return Token{Values: s}, nil
}
