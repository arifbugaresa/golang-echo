package middleware

import (
	"github.com/golang-jwt/jwt"
	"golang-echo/constants"
	"time"
)

/*
Function untuk membuat token jwt
*/
func CreateToken(userId int, name string) (string, error) {

	// Set Payload JWT
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// Set Algoritma JWT kita pake HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate bagian Signature JWT (untuk validasi jwt)
	return token.SignedString([]byte(constants.SECRET_JWT))
}
