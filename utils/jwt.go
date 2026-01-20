package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type CustomClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, username string) (string, error) {

	claims := CustomClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)), //2小时过期
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(viper.Get("JWT_SECRET"))

}

func ParseToken(tokenString string) (*CustomClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {

		return []byte(viper.GetString("JWT_SECRET")), nil
	})

	if err != nil {
		logrus.Error("token 解析失败", err)
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {

		return claims, nil
	}

	return nil, errors.New("invalid token")
}
