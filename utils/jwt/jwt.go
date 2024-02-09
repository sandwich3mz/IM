package jwt

import (
	"IM/configs"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const TokenExpireDuration = 24 * time.Hour

type Claims struct {
	UserID int64
	jwt.StandardClaims
}

var Secret = []byte(configs.Conf.JwtConfig.Secret)

// GenerateJwt 发放 Token
func GenerateJwt(userID int64) (tokenString string, err error) {
	expireTime := TokenExpireDuration
	claims := Claims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expireTime).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(Secret)
}

// ParseToken 解析 Token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("INVALID TOKEN")
}
