package util

import (
	"admin_app/global"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(account, password string) (string, error) {
	var jwtSecret = []byte(global.Config.GetString("Auth.AccessSecret"))
	var jwtExpire = global.Config.GetInt64("Auth.AccessExpire")
	nowTime := time.Now()

	claims := Claims{
		account,
		password,
		jwt.StandardClaims{
			ExpiresAt: nowTime.Unix() + jwtExpire,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	var jwtSecret = []byte(global.Config.GetString("Auth.AccessSecret"))
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
