package jwt

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type jwtHandler struct{}

func NewJWTHandler(e *echo.Echo) {
	h := &jwtHandler{}

	e.GET("/jwt", h.NewJWTToken)
}

func (h *jwtHandler) NewJWTToken(c echo.Context) error {
	jwtTokenSeed := []byte(os.Getenv("JWT_TOKEN_SEED"))

	token := jwt.New(jwt.SigningMethodHS256)

	tokenSigned, _ := token.SignedString(jwtTokenSeed)

	return c.JSON(http.StatusOK, echo.Map{"token": tokenSigned})
}
