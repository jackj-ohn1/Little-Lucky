package util

import (
	"github.com/dgrijalva/jwt-go"
	"temp/config"
	"time"
)

type Claims struct {
	Account string
	jwt.StandardClaims
}

func Generate(account string) (string, error) {
	claims := Claims{
		Account: account,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    config.Conf.Token.Issuer,
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.Conf.Token.Issuer))
}

func Parse(token string) (*Claims, error) {
	claims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Conf.Token.Issuer), nil
	})
	if err != nil {
		return nil, err
	}
	
	if claims != nil {
		if tokenClaims, ok := claims.Claims.(*Claims); ok && claims.Valid {
			return tokenClaims, nil
		}
	}
	return nil, err
}
