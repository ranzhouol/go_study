package jwtV1

import "github.com/golang-jwt/jwt/v4"

// 使用hs256算法, key字符串签名

type HS struct {
	Key string
}

// 签名
func (hs *HS) Encode(claims jwt.Claims) (string, error) {
	// 1. 创建token，传入签名方法和声明claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 2. 生成签名, 传入key的字节数组
	sign, err := token.SignedString([]byte(hs.Key))
	return sign, err
}

// 验签
func (hs *HS) Decode(sign string, claims jwt.Claims) error {
	_, err := jwt.ParseWithClaims(sign, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(hs.Key), nil
	})

	return err
}
