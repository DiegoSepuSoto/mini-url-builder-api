package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		certificateBytes := []byte(os.Getenv("JWT_TOKEN_SEED"))

		authHeader := c.Request().Header.Get("Authorization")
		if strings.Contains(authHeader, "Bearer ") {
			authHeader = strings.Split(authHeader, "Bearer ")[1]
		}

		token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
			return certificateBytes, nil
		})
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"message": err.Error()})
		}

		if !token.Valid {
			return c.JSON(http.StatusUnauthorized, echo.Map{"message": "invalid token"})
		}

		return next(c)
	}
}
