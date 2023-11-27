package jwtV1

import "github.com/golang-jwt/jwt/v4"

// 使用私钥、公钥签名

type RS struct {
	PublicKey  string
	PrivateKey string
}

// 签名
func (rs *RS) Encode(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// 解析私钥
	pKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(rs.PrivateKey))
	if err != nil {
		return "", err
	}

	// 利用私钥签名
	sign, err := token.SignedString(pKey)
	return sign, err
}

// 验签
func (rs *RS) Decode(sign string, claims jwt.Claims) error {
	_, err := jwt.ParseWithClaims(sign, claims, func(token *jwt.Token) (interface{}, error) {
		return jwt.ParseRSAPublicKeyFromPEM([]byte(rs.PublicKey))
	})

	return err
}
