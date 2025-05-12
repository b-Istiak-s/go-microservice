package middleware

import (
	"crud-service/internal/auth"
	"crud-service/internal/util/response"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func init() {
	// if os.Getenv("ENV") == "development" {
	// 	err := godotenv.Load()
	// 	if err != nil {
	// 		log.Fatal("Error loading .env file")
	// 	}
	// }
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(c, http.StatusUnauthorized, "Authorization header missing")
			c.Abort()
			return
		}

		// Expecting "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Error(c, http.StatusUnauthorized, "Invalid authorization format")
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Validate token
		claims, err := auth.ParseJWT(tokenString)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "Invalid token", err.Error())
			c.Abort()
			return
		}

		// Store user info in context for handlers to use
		c.Set("userID", claims.UserID)

		// Call the auth-service /verify endpoint
		req, err := http.NewRequest("POST", os.Getenv("LOADBALANCER_URL")+"/api/auth/verify", nil)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, "Internal server error")
			c.Abort()
			return
		}
		req.Header.Set("Authorization", authHeader)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, "Auth service unreachable")
			c.Abort()
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			response.Error(c, http.StatusUnauthorized, "User verification failed")
			c.Abort()
			return
		}

		c.Next()
	}
}
