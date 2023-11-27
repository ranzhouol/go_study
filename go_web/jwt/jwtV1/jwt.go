package jwtV1

import "github.com/golang-jwt/jwt/v4"

type JwtValidator interface {
	// 签名, jwt.Claims接口只有一个验证方法 Valid()
	Encode(claims jwt.Claims) (string, error)
	// 验签
	Decode(sign string, claims jwt.Claims) error
}
