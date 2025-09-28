package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/peetwerapat/hospital-system-api/pkg/myJwt"
	"github.com/peetwerapat/hospital-system-api/pkg/response"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, response.BaseHttpResponse{
				StatusCode: http.StatusUnauthorized,
				Message:    "Unauthorized",
			})
			c.Abort()
			return
		}

		var tokenString string
		if strings.HasPrefix(strings.ToLower(authHeader), "bearer ") {
			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 {
				c.JSON(http.StatusUnauthorized, response.BaseHttpResponse{
					StatusCode: http.StatusUnauthorized,
					Message:    "Invalid authorization header format",
				})
				c.Abort()
				return
			}
			tokenString = parts[1]
		} else {
			tokenString = authHeader
		}

		claims := &myJwt.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(myJwt.GetJWTSecret()), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, response.BaseHttpResponse{
				StatusCode: http.StatusUnauthorized,
				Message:    "Invalid token",
			})
			c.Abort()
			return
		}

		c.Set("staff_id", claims.ID)
		c.Set("hospital_id", claims.HospitalID)
		c.Next()
	}
}
