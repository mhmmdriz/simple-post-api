package middlewares

import (
	"fmt"
	"soal-eksplorasi/constants"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	UserId   int    `json:"userId"`
	UserName string `json:"userName"`
	IsAdmin  bool   `json:"isAdmin"`
	jwt.RegisteredClaims
}

func GenerateTokenJWT(userId int, userName string, isAdmin bool) string {
	fmt.Println(userId)
	var userClaims = JwtCustomClaims{
		userId, userName, isAdmin,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)

	resultJWT, _ := token.SignedString([]byte(constants.PRIVATE_KEY_JWT))
	return resultJWT
}
