package middleware

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte("G4C0R"),
	})
}

func CreateToken(userId int, email string, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["Authorization"] = true
	claims["userId"] = userId
	claims["email"] = email
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("G4C0R"))
}

func ExtractToken(c echo.Context) int {
	headerData := c.Request().Header.Get("Authorization")
	dataAuth := strings.Split(headerData, " ")
	token := dataAuth[len(dataAuth)-1]
	datajwt, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("G4C0R"), nil
	})

	if datajwt.Valid {
		claims := datajwt.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return int(userId)
	}

	return -1
}