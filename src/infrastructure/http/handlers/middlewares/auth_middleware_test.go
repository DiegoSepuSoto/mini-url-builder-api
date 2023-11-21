package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func testHandler(c echo.Context) error {
	return c.String(http.StatusOK, "response ok")
}

func TestAuthMiddleware(t *testing.T) {
	t.Run("when auth middleware accepts a request with correct token", func(t *testing.T) {
		t.Setenv("JWT_TOKEN_SEED", "supersecret")

		e := echo.New()
		e.Use(AuthMiddleware)
		e.GET("/", testHandler)
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.8x2hIBGylPBtKnAoEP8wJqqXbXaQyOK0z8bjpasZGfo")

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "response ok", rec.Body.String())
	})

	t.Run("when auth middleware rejects a request with invalid/no token", func(t *testing.T) {
		t.Setenv("JWT_TOKEN_SEED", "supersecret")

		e := echo.New()
		e.Use(AuthMiddleware)
		e.GET("/", testHandler)
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Equal(t, "{\"message\":\"token contains an invalid number of segments\"}\n", rec.Body.String())
	})
}
