package main

import (
	jwt "github.com/golang-jwt/jwt/v5"
	"time"
)

type JWT struct {
	SigningKey []byte
}

// Claims jwt中自定义信息，可以自己添加各种信息，比如username、userid等用户相关信息
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte("test")

func GenerateToken(username string) (string, error) {
	// 创建playload
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    username,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ParseToken 解析token
func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	return claims, err
}
