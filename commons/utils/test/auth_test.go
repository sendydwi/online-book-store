package utils_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/sendydwi/online-book-store/commons/utils"
	"github.com/stretchr/testify/assert"
)

func TestCheckAuth(t *testing.T) {
	_ = godotenv.Load(".env")
	gin.SetMode(gin.TestMode)

	// Helper function to create a test context
	createTestContext := func(authHeader string) (*gin.Context, *httptest.ResponseRecorder) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		if authHeader != "" {
			req.Header.Set("Authorization", authHeader)
		}
		c.Request = req

		return c, w
	}

	t.Run("missing authorization header", func(t *testing.T) {
		c, w := createTestContext("")
		utils.CheckAuth(c)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "Authorization header is missing")
	})

	t.Run("invalid_token_format", func(t *testing.T) {
		c, w := createTestContext("BearerToken")
		utils.CheckAuth(c)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "Invalid token format")
	})

	t.Run("invalid_token", func(t *testing.T) {
		c, w := createTestContext("Bearer invalid.token.here")
		utils.CheckAuth(c)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "Invalid or expired token")
	})

	t.Run("expired_token", func(t *testing.T) {
		// Create an expired token
		claims := jwt.MapClaims{
			"exp": time.Now().Add(-time.Hour).Unix(),
			"id":  "123",
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, _ := token.SignedString([]byte(os.Getenv("SECRET")))

		c, w := createTestContext("Bearer " + tokenString)
		utils.CheckAuth(c)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "Invalid or expired token")
	})

	t.Run("valid token", func(t *testing.T) {
		claims := jwt.MapClaims{
			"exp": time.Now().Add(time.Hour).Unix(),
			"id":  "123",
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, _ := token.SignedString([]byte(os.Getenv("SECRET")))

		c, w := createTestContext("Bearer " + tokenString)
		utils.CheckAuth(c)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "123", c.GetString("userId"))
	})
}
